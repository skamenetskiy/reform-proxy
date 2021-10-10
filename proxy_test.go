package proxy

import (
	"database/sql"
	"reflect"
	"testing"

	"gopkg.in/reform.v1"
)

func TestNew(t *testing.T) {
	type args struct {
		primary     DB
		secondaries []DB
	}
	tests := []struct {
		name string
		args args
		want Proxy
	}{
		{
			"primary only",
			args{
				primary: reform.NewDB(nil, nil, nil),
			},
			&proxy{
				primary: reform.NewDB(nil, nil, nil),
			},
		},
		{
			"one secondary",
			args{
				primary: reform.NewDB(nil, nil, nil),
				secondaries: []DB{
					reform.NewDB(nil, nil, nil),
				},
			},
			&proxy{
				primary: reform.NewDB(nil, nil, nil),
				secondaries: []DB{
					reform.NewDB(nil, nil, nil),
				},
			},
		},
		{
			"two secondaries",
			args{
				primary: reform.NewDB(nil, nil, nil),
				secondaries: []DB{
					reform.NewDB(nil, nil, nil),
					reform.NewDB(nil, nil, nil),
				},
			},
			&proxy{
				primary: reform.NewDB(nil, nil, nil),
				secondaries: []DB{
					reform.NewDB(nil, nil, nil),
					reform.NewDB(nil, nil, nil),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.primary, tt.args.secondaries...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_Count(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view reform.View
		tail string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.Count(tt.args.view, tt.args.tail, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_Delete(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		record reform.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.Delete(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_DeleteFrom(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view reform.View
		tail string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.DeleteFrom(tt.args.view, tt.args.tail, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_FindAllFrom(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view   reform.View
		column string
		args   []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []reform.Struct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.FindAllFrom(tt.args.view, tt.args.column, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAllFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_FindByPrimaryKeyFrom(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		table reform.Table
		pk    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    reform.Record
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.FindByPrimaryKeyFrom(tt.args.table, tt.args.pk)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByPrimaryKeyFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByPrimaryKeyFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_FindByPrimaryKeyTo(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		record reform.Record
		pk     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.FindByPrimaryKeyTo(tt.args.record, tt.args.pk); (err != nil) != tt.wantErr {
				t.Errorf("FindByPrimaryKeyTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_FindOneFrom(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view   reform.View
		column string
		arg    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    reform.Struct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.FindOneFrom(tt.args.view, tt.args.column, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOneFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOneFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_FindOneTo(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		str    reform.Struct
		column string
		arg    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.FindOneTo(tt.args.str, tt.args.column, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("FindOneTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_FindRows(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view   reform.View
		column string
		arg    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.FindRows(tt.args.view, tt.args.column, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_Insert(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		str reform.Struct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.Insert(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_InsertColumns(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		str     reform.Struct
		columns []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.InsertColumns(tt.args.str, tt.args.columns...); (err != nil) != tt.wantErr {
				t.Errorf("InsertColumns() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_InsertMulti(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		structs []reform.Struct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.InsertMulti(tt.args.structs...); (err != nil) != tt.wantErr {
				t.Errorf("InsertMulti() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_Primary(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if got := p.Primary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Primary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_Save(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		record reform.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.Save(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_Secondaries(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   []DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if got := p.Secondaries(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Secondaries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_Secondary(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if got := p.Secondary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Secondary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_SelectAllFrom(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view reform.View
		tail string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []reform.Struct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.SelectAllFrom(tt.args.view, tt.args.tail, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectAllFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectAllFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_SelectOneFrom(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view reform.View
		tail string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    reform.Struct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.SelectOneFrom(tt.args.view, tt.args.tail, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectOneFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectOneFrom() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_SelectOneTo(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		str  reform.Struct
		tail string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.SelectOneTo(tt.args.str, tt.args.tail, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("SelectOneTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_SelectRows(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		view reform.View
		tail string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.SelectRows(tt.args.view, tt.args.tail, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proxy_Update(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		record reform.Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.Update(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_UpdateColumns(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		record  reform.Record
		columns []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			if err := p.UpdateColumns(tt.args.record, tt.args.columns...); (err != nil) != tt.wantErr {
				t.Errorf("UpdateColumns() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_proxy_UpdateView(t *testing.T) {
	type fields struct {
		primary     DB
		secondaries []DB
		next        uint32
	}
	type args struct {
		str     reform.Struct
		columns []string
		tail    string
		args    []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &proxy{
				primary:     tt.fields.primary,
				secondaries: tt.fields.secondaries,
				next:        tt.fields.next,
			}
			got, err := p.UpdateView(tt.args.str, tt.args.columns, tt.args.tail, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateView() got = %v, want %v", got, tt.want)
			}
		})
	}
}
