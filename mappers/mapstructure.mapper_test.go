package mappers

import (
	"context"
	"testing"
)

type UserDTO struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserEntity struct {
	ID string `bson:"_id"`
	Name string `bson:"name"`
}

func TestMapstructureMapper(t *testing.T) {
	mapper := NewMapStructureMapper[UserDTO, UserDTO, UserDTO, UserEntity, UserEntity, UserEntity]()

	c := context.TODO()

	userDto := &UserDTO{
		ID: "1",
		Name: "张三",
	}

	userEntity, err := mapper.ConvertToEntity(c, userDto)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("userEntity: %v\n", userEntity)

	userEntities, err := mapper.ConvertToEntities(c, []*UserDTO{userDto})
	if err != nil {
		t.Fatal(err)
	}
	for _, entity := range userEntities {
		t.Logf("userEntity: %v\n", entity)
	}
}
