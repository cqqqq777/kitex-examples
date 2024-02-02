/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int32  `gorm:"column:id;primaryKey" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Age       int32  `gorm:"column:age" json:"age"`
	Gender    string `gorm:"column:gender" json:"gender"`
	Introduce string `gorm:"column:introduce" json:"introduce"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}