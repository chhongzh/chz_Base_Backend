package action

import "github.com/gin-gonic/gin"

type MetaChain struct {
	MetaFields map[string]any
}

func NewMetaChain() *MetaChain {
	return &MetaChain{
		MetaFields: make(map[string]any),
	}
}

func (m *MetaChain) Add(key string, value any) *MetaChain {
	m.MetaFields[key] = value

	return m
}

func (m *MetaChain) WithClientInfo(c *gin.Context) *MetaChain {
	return m.Add("Remote IP", c.ClientIP()).
		Add("User Agent", c.Request.UserAgent())
}

func (m *MetaChain) Remove(key string) *MetaChain {
	delete(m.MetaFields, key)

	return m
}

func (m *MetaChain) Build() map[string]any {
	return m.MetaFields
}
