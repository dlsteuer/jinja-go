package nodes

import "fmt"

type CommentNode struct {
	HasNoChildren
	body string
}

func NewCommentNode(body string) INode {
	return &CommentNode{NewHasNoChildren("COMMENT", body), body}
}

func (node *CommentNode) ToString() string {
	return fmt.Sprintf("<COMMENT>")
}

func (node *CommentNode) Type() int {
	return NODE_TYPE_COMMENT
}

func (node *CommentNode) Render(c map[string]interface{}) string {
	return ""
}
