// Copyright (c) 2012 - Cloud Instruments Co. Ltd.
// 
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package seelog

import (
	"testing"
	"strconv"
	"os"
)

func Test_Asyncloop(t *testing.T) {
	switchToRealFSWrapper(t)
	fileName := "log.log"
	count := 100
	
	os.Remove(fileName)
	
	testConfig := `
<seelog type="asyncloop">
	<outputs formatid="msg">
		<file path="` + fileName + `"/>
	</outputs>
	<formats>
		<format id="msg" format="%Msg%n"/>
	</formats>
</seelog>`

	logger, _ := LoggerFromConfigAsBytes([]byte(testConfig))
	err := ReplaceLogger(logger)
	if err != nil {
		t.Error(err)
		return
	}
	
	for i := 0; i < count; i++ {
		Trace(strconv.Itoa(i))
	}
	
	Flush()
	
	gotCount, err := countSequencedRowsInFile(fileName)
	if err != nil {
		t.Error(err)
		return
	}
	
	if int64(count) != gotCount {
		t.Errorf("Wrong count of log messages. Expected: %v, got: %v.", count, gotCount)
		return
	}
}

func Test_AsyncloopOff(t *testing.T) {
	switchToRealFSWrapper(t)
	fileName := "log.log"
	count := 100
	
	os.Remove(fileName)
	
	testConfig := `
<seelog type="asyncloop" levels="off">
	<outputs formatid="msg">
		<file path="` + fileName + `"/>
	</outputs>
	<formats>
		<format id="msg" format="%Msg%n"/>
	</formats>
</seelog>`

	logger, _ := LoggerFromConfigAsBytes([]byte(testConfig))
	err := ReplaceLogger(logger)
	if err != nil {
		t.Error(err)
		return
	}
	
	for i := 0; i < count; i++ {
		Trace(strconv.Itoa(i))
	}
	
	Flush()
	
	gotCount, err := countSequencedRowsInFile(fileName)
	if err != nil {
		t.Error(err)
		return
	}
	
	if gotCount != 0 {
		t.Errorf("Wrong count of log messages. Expected: %v, got: %v.", 0, gotCount)
		return
	}
}
