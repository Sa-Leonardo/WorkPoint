package models


import (
    "time"
)

// Empresa representa cada cliente (multiempresa)
type Empresa struct {
    ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Nome     string    `gorm:"size:100;not null" json:"nome"`
    CNPJ     string    `gorm:"size:18;unique" json:"cnpj"`
    Email    string    `gorm:"size:120" json:"email"`
    Telefone string    `gorm:"size:20" json:"telefone"`
    CriadoEm time.Time `gorm:"autoCreateTime" json:"criado_em"`

    Usuarios []Usuario `gorm:"foreignKey:EmpresaID" json:"usuarios,omitempty"`
}

// Usuario representa cada funcionário ou usuário do sistema
type Usuario struct {
    ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    EmpresaID uint      `gorm:"not null" json:"empresa_id"`
    Nome      string    `gorm:"size:100;not null" json:"nome"`
    Email     string    `gorm:"size:120;unique;not null" json:"email"`
    Role      string    `gorm:"size:50;default:'usuario'" json:"role"` // funcionario, gestor, admin
    Ativo     bool      `gorm:"default:true" json:"ativo"`
    CriadoEm  time.Time `gorm:"autoCreateTime" json:"criado_em"`

    Registros      []RegistroPonto `gorm:"foreignKey:UsuarioID" json:"registros,omitempty"`
    Justificativas []Justificativa `gorm:"foreignKey:UsuarioID" json:"justificativas,omitempty"`
}

// RegistroPonto representa cada batida de ponto do usuário
type RegistroPonto struct {
    ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    UsuarioID uint       `gorm:"not null" json:"usuario_id"`
    Data      time.Time  `gorm:"not null" json:"data"`
    Entrada   *time.Time `json:"entrada,omitempty"`
    Saida     *time.Time `json:"saida,omitempty"`
    CriadoEm  time.Time  `gorm:"autoCreateTime" json:"criado_em"`

    Justificativas []Justificativa `gorm:"foreignKey:RegistroID" json:"justificativas,omitempty"`
}

// Justificativa serve para faltas ou atrasos
type Justificativa struct {
    ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    UsuarioID  uint      `gorm:"not null" json:"usuario_id"`
    RegistroID uint      `gorm:"not null" json:"registro_id"`
    Motivo     string    `gorm:"type:text" json:"motivo"`
    Aprovado   bool      `gorm:"default:false" json:"aprovado"`
    CriadoEm   time.Time `gorm:"autoCreateTime" json:"criado_em"`
}
