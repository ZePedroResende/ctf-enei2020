package main

import (
	"bufio"
	"encoding/base64"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"

	"fmt"
	"github.com/soheilhy/cmux"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "github.com/ZePedroResende/ctf-enei2020/6.XSS/flag"
)

type FlagServer struct {
	pb.UnimplementedFlagServer
}

func httpServe(l net.Listener) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkhealth/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	s := &http.Server{Handler: mux}
	return s.Serve(l)
}
func (*FlagServer) getFlag(ctx context.Context, in *pb.FlagRequest) (*pb.FlagReply, error) {

	imgFile, err := os.Open("static/flag.png") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	if in.Pass != "" {
		return &pb.FlagReply{Message: imgBase64Str}, nil
	}
	return &pb.FlagReply{Message: "failed"}, nil
}

func grpcServe(l net.Listener) error {
	s := grpc.NewServer()
	pb.RegisterFlagServer(s, &FlagServer{})

	return s.Serve(l)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(grpcListener) })
	g.Go(func() error { return httpServe(httpListener) })
	g.Go(func() error { return m.Serve() })

	log.Println("run server:", g.Wait())
}
