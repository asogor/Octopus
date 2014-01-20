package octopus

import (
	"testing"
	"regexp"
	"log"
	"strings"
	"time"
)

func TestContextNil(t *testing.T) {
	re1, _ := regexp.Compile(`[Le]{3}`)
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}else{
			t.Fail()
		}
	}()
	NewFunnel(1,re1,100,nil,nil);
}

func TestGotResult(t *testing.T) {
	re1, _ := regexp.Compile(`[Le]{2}`)
	result := &CounterAction{1,0,0}

	m := NewFunnel(1,re1,100,result,&Context{})
	m.GetCurrentState()

	reader := strings.NewReader(`Mrs. Leonora Spocky`)
	var r rune = 0
	var size int = 0;
	var err error = nil
	for(err == nil){
		r,size,err = reader.ReadRune()
		log.Printf("Write %#v %#v %#v",r,size,err)
		m.WriteRune(r,size,err)
	}

	time.Sleep(10 * time.Millisecond)

	if!((result.CompleteCounter == 1)&&(result.MatchCounter == 1)){
		t.Fatal("Failed Result %#v",result);
	}
}


func TestBuildGroup(t *testing.T) {
	reg1, _ := regexp.Compile(`[Le]{2}`)
	result1 := &CounterAction{1,0,0}
	reg2, _ := regexp.Compile(`[Spo]{3}`)
	result2 := &CounterAction{2,0,0}
	context := &Context{"session1","path1"}

	fg := NewFunnelGroup(context)
	fg.AddFunnel(1,reg1,100,result1)
	fg.AddFunnel(2,reg2,100,result2)

	reader := strings.NewReader(`Mrs. Leonora Spocky`)
	var r rune = 0
	var size int = 0;
	var err error = nil
	for(err == nil){
		r,size,err = reader.ReadRune()
		log.Printf("Write %#v %#v %#v",r,size,err)
		fg.WriteRune(r,size,err)
	}

	time.Sleep(10 * time.Millisecond)

	if!((result1.CompleteCounter == 1)&&(result1.MatchCounter == 1)){
		t.Fatal("Failed Result %#v",result1);
	}

	if!((result2.CompleteCounter == 1)&&(result2.MatchCounter == 1)){
		t.Fatal("Failed Result %#v",result2);
	}
}
