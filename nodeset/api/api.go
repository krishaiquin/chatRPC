package api

type Node struct {
	NodeId   uint32
	UserName string
	Addr     string
}

type AddArgs struct {
	Addr     string
	Username string
}

type AddRet struct {
	NodeId  uint32
	NodeSet []Node
}
