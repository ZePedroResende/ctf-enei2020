from flask import Flask, request, render_template
from xml.dom import minidom

app = Flask(__name__)
app.config['DEBUG'] = True

USERNAME = 'admin' # 账号
PASSWORD = 'admin' # 密码

@app.route("/")
def home():
    return render_template("index.html")

@app.route("/doLogin", methods=['POST', 'GET'])
def doLogin():
    result = None
    try:
        DOMTree = minidom.parseString(request.data)
        username = DOMTree.getElementsByTagName("username")
        username = username[0].childNodes[0].nodeValue
        password = DOMTree.getElementsByTagName("password")
        password = password[0].childNodes[0].nodeValue

        if username == USERNAME and password == PASSWORD:
            result = "<result><code>%d</code><msg>%s</msg></result>" % (1,username)
        else:
            result = "<result><code>%d</code><msg>%s</msg></result>" % (0,username)
    except Exception as e:
        result = "<result><code>%d</code><msg>%s</msg></result>" % (3,e.message)

    return result,{'Content-Type': 'text/xml;charset=UTF-8'}

if __name__ == "__main__":
    app.run()
