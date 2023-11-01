package bufio_demo

import "fmt"

func (p *BufioDemo) write() {
	data := "this writer tests1.."
	n, e := p.BfWriter.Write([]byte(data))
	if e != nil {
		fmt.Println("op write fail, e: ", e)
		return 
	}
	if n != len(data) {
		fmt.Println("op data len not ret")
		return 
	}

	p.BfWriter.Flush() // after Write()， need to call Flush() .
	fmt.Println("\nwrite to stdout succ ...")
}

func (p *BufioDemo) writeString() {
	data := "this writer tests2..."
	n, e := p.BfWriter.WriteString((data))
	if e != nil {
		fmt.Println("op write fail, e: ", e)
		return 
	}
	if n != len(data) {
		fmt.Println("op data len not ret")
		return 
	}
	fmt.Println()
	p.BfWriter.Flush() // after Write()， need to call Flush() .
	fmt.Println("\na)---> write string succ ...")
}