package swag

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPropertyNameSelectorExpr(t *testing.T) {
	input := &ast.SelectorExpr{
		X: &ast.Ident{
			NamePos: 1136,
			Name:    "time",
			Obj:     (*ast.Object)(nil),
		},
		Sel: &ast.Ident{
			NamePos: 1141,
			Name:    "Time",
			Obj:     (*ast.Object)(nil),
		},
	}
	expected := propertyName{
		STRING,
		STRING,
		"",
	}
	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameIdentObjectId(t *testing.T) {
	input := &ast.SelectorExpr{
		X: &ast.Ident{
			NamePos: 1136,
			Name:    "hoge",
			Obj:     (*ast.Object)(nil),
		},
		Sel: &ast.Ident{
			NamePos: 1141,
			Name:    "ObjectId",
			Obj:     (*ast.Object)(nil),
		},
	}
	expected := propertyName{
		STRING,
		STRING,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameIdentUUID(t *testing.T) {
	input := &ast.SelectorExpr{
		X: &ast.Ident{
			NamePos: 1136,
			Name:    "hoge",
			Obj:     (*ast.Object)(nil),
		},
		Sel: &ast.Ident{
			NamePos: 1141,
			Name:    "uuid",
			Obj:     (*ast.Object)(nil),
		},
	}
	expected := propertyName{
		STRING,
		STRING,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameIdentDecimal(t *testing.T) {
	input := &ast.SelectorExpr{
		X: &ast.Ident{
			NamePos: 1136,
			Name:    "hoge",
			Obj:     (*ast.Object)(nil),
		},
		Sel: &ast.Ident{
			NamePos: 1141,
			Name:    "Decimal",
			Obj:     (*ast.Object)(nil),
		},
	}
	expected := propertyName{
		NUMBER,
		STRING,
		"",
	}
	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameIdentTime(t *testing.T) {
	input := &ast.SelectorExpr{
		X: &ast.Ident{
			NamePos: 1136,
			Name:    "hoge",
			Obj:     (*ast.Object)(nil),
		},
		Sel: &ast.Ident{
			NamePos: 1141,
			Name:    "Time",
			Obj:     (*ast.Object)(nil),
		},
	}
	expected := propertyName{
		STRING,
		STRING,
		"",
	}

	propertyName, err := getPropertyName("test", input, nil)
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameStarExprIdent(t *testing.T) {
	input := &ast.StarExpr{
		Star: 1026,
		X: &ast.Ident{
			NamePos: 1027,
			Name:    "string",
			Obj:     (*ast.Object)(nil),
		},
	}
	expected := propertyName{
		STRING,
		STRING,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameStarExprMap(t *testing.T) {
	input := &ast.StarExpr{
		Star: 1026,
		X: &ast.MapType{
			Map: 1027,
			Key: &ast.Ident{
				NamePos: 1034,
				Name:    "string",
				Obj:     (*ast.Object)(nil),
			},
			Value: &ast.Ident{
				NamePos: 1041,
				Name:    "string",
				Obj:     (*ast.Object)(nil),
			},
		},
	}
	expected := propertyName{
		OBJECT,
		OBJECT,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameArrayStarExpr(t *testing.T) {
	input := &ast.ArrayType{
		Lbrack: 465,
		Len:    nil,
		Elt: &ast.StarExpr{
			X: &ast.Ident{
				NamePos: 467,
				Name:    "string",
				Obj:     (*ast.Object)(nil),
			},
		},
	}
	expected := propertyName{
		ARRAY,
		STRING,
		"",
	}
	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameArrayStarExprSelector(t *testing.T) {
	input := &ast.ArrayType{
		Lbrack: 1111,
		Len:    nil,
		Elt: &ast.StarExpr{
			X: &ast.SelectorExpr{
				X: &ast.Ident{
					NamePos: 1136,
					Name:    "hoge",
					Obj:     (*ast.Object)(nil),
				},
				Sel: &ast.Ident{
					NamePos: 1141,
					Name:    "ObjectId",
					Obj:     (*ast.Object)(nil),
				},
			},
		},
	}
	expected := propertyName{
		ARRAY,
		STRING,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameArrayStructType(t *testing.T) {
	input := &ast.ArrayType{
		Lbrack: 1111,
		Len:    nil,
		Elt:    &ast.StructType{},
	}
	expected := propertyName{
		ARRAY,
		OBJECT,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameMap(t *testing.T) {
	input := &ast.MapType{
		Key: &ast.Ident{
			Name: "string",
		},
		Value: &ast.Ident{
			Name: "string",
		},
	}
	expected := propertyName{
		OBJECT,
		OBJECT,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameStruct(t *testing.T) {
	input := &ast.StructType{}
	expected := propertyName{
		OBJECT,
		OBJECT,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameInterface(t *testing.T) {
	input := &ast.InterfaceType{}
	expected := propertyName{
		OBJECT,
		OBJECT,
		"",
	}

	propertyName, err := getPropertyName("test", input, New())
	assert.NoError(t, err)
	assert.Equal(t, expected, propertyName)
}

func TestGetPropertyNameChannel(t *testing.T) {
	input := &ast.ChanType{}
	_, err := getPropertyName("test", input, New())
	assert.Error(t, err)
}

func TestParseTag(t *testing.T) {
	searchDir := "testdata/tags"
	mainAPIFile := "main.go"
	p := New(SetMarkdownFileDirectory(searchDir))
	p.PropNamingStrategy = PascalCase
	err := p.ParseAPI(searchDir, mainAPIFile)
	assert.NoError(t, err)

	if len(p.swagger.Tags) != 3 {
		t.Log(len(p.swagger.Tags))
		t.Log("Number of tags did not match")
		t.FailNow()
	}

	dogs := p.swagger.Tags[0]
	if dogs.TagProps.Name != "dogs" || dogs.TagProps.Description != "Dogs are cool" {
		t.Log("Failed to parse dogs name or description")
		t.FailNow()
	}

	cats := p.swagger.Tags[1]
	if cats.TagProps.Name != "cats" || cats.TagProps.Description != "Cats are the devil" {
		t.Log("Failed to parse cats name or description")
		t.FailNow()
	}
}
