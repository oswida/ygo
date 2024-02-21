package common_test

import (
	"bytes"
	"reflect"
	"testing"
	"ygo/common"
)

func TestIdWrite(t *testing.T) {
	ident := common.NewId(1, 1)
	buf := bytes.NewBuffer([]byte{})
	err := ident.Write(buf)
	if err != nil {
		t.Errorf("Problem writing identifier %v", err)
	}
	ident = common.NewId(1024, 1024)
	err = ident.Write(buf)
	if err != nil {
		t.Errorf("Problem writing identifier %v", err)
	}
	reference := []byte{1, 1, 128, 8, 128, 8}
	if !reflect.DeepEqual(reference, buf.Bytes()) {
		t.Errorf("expect %v got %v", reference, buf.Bytes())
	}
}

func TestIdRead(t *testing.T) {
	ident := common.NewId(1, 2)
	buf := bytes.NewBuffer([]byte{})
	err := ident.Write(buf)
	if err != nil {
		t.Errorf("Problem writing identifier %v", err)
	}
	ident = common.NewId(1024, 1025)
	err = ident.Write(buf)
	if err != nil {
		t.Errorf("Problem writing identifier %v", err)
	}
	source := bytes.NewBuffer(buf.Bytes())
	firstId, err := common.GenerateId()
	if err != nil {
		t.Errorf("Problem generating identifier client %v", err)
	}
	firstId.Read(source)
	reference := common.NewId(1, 2)
	if !reflect.DeepEqual(reference, firstId) {
		t.Errorf("expect %v got %v", reference, firstId)
	}
	secondId, err := common.GenerateId()
	if err != nil {
		t.Errorf("Problem generating identifier client %v", err)
	}
	err = secondId.Read(source)
	if err != nil {
		t.Errorf("Problem reading identifier %v", err)
	}
	reference = common.NewId(1024, 1025)
	if !reflect.DeepEqual(reference, secondId) {
		t.Errorf("expect %v got %v", reference, secondId)
	}
}
