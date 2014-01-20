/**
 */
package octopus

import (
	"io"
	"fmt"
)

type DebugReader struct {
	origin io.RuneReader
}

func NewDebugReader(o io.RuneReader) (result DebugReader) {
	return DebugReader{origin:o}
}

func (dr DebugReader) ReadRune() (r rune, size int, err error){
	r,size,err = dr.origin.ReadRune()
	fmt.Printf("Read %d %d %d \n",r,size,err)
	return r,size,err
}
