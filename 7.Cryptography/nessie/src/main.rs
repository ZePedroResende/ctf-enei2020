extern crate crypto;

use self::crypto::digest::Digest;
use self::crypto::sha1::Sha1;
use obfstr;
use std::env;
use std::panic;
use std::time::SystemTime;
use ultra::Enigma;

fn hour() -> u8 {
    let tnow = SystemTime::now()
        .duration_since(SystemTime::UNIX_EPOCH)
        .unwrap()
        .as_secs();
    // 60*60*24 = 86400
    // 60*60    = 1440
    // With day savings
    //let hours = ((tnow % 86400) / 3600) + 1;
    // Without day savings
    let hours = (tnow % 86400) / 3600;
    return hours as u8;
}

fn main() {
    panic::set_hook(Box::new(|_| {
        println!("Nice try.");
    }));

    let no_access = obfstr::obflocal!("No Access.");
    let lapras = obfstr::obflocal!("Lapras");
    let url = obfstr::obflocal!("https://upload.wikimedia.org/wikipedia/en/thumb/8/80/Wikipedia-logo-v2.svg/1024px-Wikipedia-logo-v2.svg.png");

    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("Needs the key!");
        return;
    }

    let h = hour();
    if h > 6 && h < 20 {
        println!("Loch Ness Monster is not here.");
        return;
    }

    if &args[1] == "help" {
        println!("Think Transport Pokémon. Think git commit.");
    }

    let hashkey = &args[1];
    let mut hasher = Sha1::new();
    hasher.input_str(lapras.as_str());
    let hex = hasher.result_str();

    if hashkey == hex.as_str() {
        let rotors = "131";
        let key = "NES";
        let ring = "SIE";
        let mut enigma = Enigma::new(&rotors, &key, &ring, 'B', "");
        println!("Solve the enigma.");
        println!("{}", enigma.encrypt(url.as_str()));
        println!("{} {}{}", &rotors, &key, &ring);
    } else {
        println!("{}", no_access.as_str());
    }
}
