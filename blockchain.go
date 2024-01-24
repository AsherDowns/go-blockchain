package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data        map[string]interfact{}
	hash        string	
	prevHash    string
	timestamp   time.Time
	powchance   int
}