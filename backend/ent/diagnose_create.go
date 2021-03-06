// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/ADMIN/app/ent/department"
	"github.com/ADMIN/app/ent/diagnose"
	"github.com/ADMIN/app/ent/disease"
	"github.com/ADMIN/app/ent/doctor"
	"github.com/ADMIN/app/ent/patient"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DiagnoseCreate is the builder for creating a Diagnose entity.
type DiagnoseCreate struct {
	config
	mutation *DiagnoseMutation
	hooks    []Hook
}

// SetDiseaseID sets the disease edge to Disease by id.
func (dc *DiagnoseCreate) SetDiseaseID(id int) *DiagnoseCreate {
	dc.mutation.SetDiseaseID(id)
	return dc
}

// SetNillableDiseaseID sets the disease edge to Disease by id if the given value is not nil.
func (dc *DiagnoseCreate) SetNillableDiseaseID(id *int) *DiagnoseCreate {
	if id != nil {
		dc = dc.SetDiseaseID(*id)
	}
	return dc
}

// SetDisease sets the disease edge to Disease.
func (dc *DiagnoseCreate) SetDisease(d *Disease) *DiagnoseCreate {
	return dc.SetDiseaseID(d.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (dc *DiagnoseCreate) SetDepartmentID(id int) *DiagnoseCreate {
	dc.mutation.SetDepartmentID(id)
	return dc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (dc *DiagnoseCreate) SetNillableDepartmentID(id *int) *DiagnoseCreate {
	if id != nil {
		dc = dc.SetDepartmentID(*id)
	}
	return dc
}

// SetDepartment sets the department edge to Department.
func (dc *DiagnoseCreate) SetDepartment(d *Department) *DiagnoseCreate {
	return dc.SetDepartmentID(d.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (dc *DiagnoseCreate) SetPatientID(id int) *DiagnoseCreate {
	dc.mutation.SetPatientID(id)
	return dc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (dc *DiagnoseCreate) SetNillablePatientID(id *int) *DiagnoseCreate {
	if id != nil {
		dc = dc.SetPatientID(*id)
	}
	return dc
}

// SetPatient sets the patient edge to Patient.
func (dc *DiagnoseCreate) SetPatient(p *Patient) *DiagnoseCreate {
	return dc.SetPatientID(p.ID)
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (dc *DiagnoseCreate) SetDoctorID(id int) *DiagnoseCreate {
	dc.mutation.SetDoctorID(id)
	return dc
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (dc *DiagnoseCreate) SetNillableDoctorID(id *int) *DiagnoseCreate {
	if id != nil {
		dc = dc.SetDoctorID(*id)
	}
	return dc
}

// SetDoctor sets the doctor edge to Doctor.
func (dc *DiagnoseCreate) SetDoctor(d *Doctor) *DiagnoseCreate {
	return dc.SetDoctorID(d.ID)
}

// Mutation returns the DiagnoseMutation object of the builder.
func (dc *DiagnoseCreate) Mutation() *DiagnoseMutation {
	return dc.mutation
}

// Save creates the Diagnose in the database.
func (dc *DiagnoseCreate) Save(ctx context.Context) (*Diagnose, error) {
	var (
		err  error
		node *Diagnose
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiagnoseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiagnoseCreate) SaveX(ctx context.Context) *Diagnose {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DiagnoseCreate) sqlSave(ctx context.Context) (*Diagnose, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DiagnoseCreate) createSpec() (*Diagnose, *sqlgraph.CreateSpec) {
	var (
		d     = &Diagnose{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: diagnose.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diagnose.FieldID,
			},
		}
	)
	if nodes := dc.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diagnose.DiseaseTable,
			Columns: []string{diagnose.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diagnose.DepartmentTable,
			Columns: []string{diagnose.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diagnose.PatientTable,
			Columns: []string{diagnose.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   diagnose.DoctorTable,
			Columns: []string{diagnose.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
