package neurogo

// Related packages and documentation:
// https://pkg.go.dev/github.com/oschwald/maxminddb-golang#example-Reader.Lookup-Interface
// https://pkg.go.dev/encoding/binary#example-Read-Multi
// maybe https://www.jonathan-petitcolas.com/2014/09/25/parsing-binary-files-in-go.html, but it's old
//
// https://github.com/dfsp-spirit/libfs/blob/main/include/libfs.h#L2023 for the fs surface file format

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func read_fs_mesh(filepath string) {
	//b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	//r := bytes.NewReader(b)

	file, err := os.Open(filepath)
	if err != nil {
   		panic(err)
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
	   fmt.Println(err)
	   return
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
	   fmt.Println(err)
	   return
	}

	// Read the byte slice
	r := bytes.NewReader(bs)

	var header_part1 struct {
		magic_b1 uint8
		magic_b2 uint8
		magic_b3 uint8
		//Mine [3]byte
	}

	if err := binary.Read(r, binary.LittleEndian, &header_part1); err != nil {
		fmt.Println("binary.Read failed on first part of fs surface header:", err)
	}

	fmt.Println(header_part1.magic_b1)
	fmt.Println(header_part1.magic_b2)
	fmt.Println(header_part1.magic_b3)

	createdLine, err := common.readNewlineTerminatedString(r);
    commentLine, err := common.readNewlineTerminatedString(r);

	var header_part2 struct {
		num_verts int32
		num_faces int32
	}

	if err := binary.Read(r, binary.LittleEndian, &header_part2); err != nil {
		fmt.Println("binary.Read failed on second part of fs surface header:", err)
	}

	fmt.Println(header_part2.num_verts)
	fmt.Println(header_part2.num_faces)
}