package snowflake

import "github.com/bwmarrin/snowflake"

type Snowflake interface {
	Generate() snowflake.ID
}

type snowflakeProvider struct {
	node *snowflake.Node
}

func (s *snowflakeProvider) Generate() snowflake.ID {
	return s.node.Generate()
}

func NewSnowflake(node int64) (Snowflake, error) {
	var err error
	u := new(snowflakeProvider)
	snowflake.Epoch = 1640995200000
	u.node, err = snowflake.NewNode(node)
	return u.node, err
}
