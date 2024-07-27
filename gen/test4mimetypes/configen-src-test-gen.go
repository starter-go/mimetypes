package test4mimetypes
import (
    p85a4d026d "github.com/starter-go/mimetypes"
    p37b02ea1a "github.com/starter-go/mimetypes/src/test/golang/unit"
    p277ab6ae7 "github.com/starter-go/mimetypes/src/test/golang/unit/demotypes"
     "github.com/starter-go/application"
)

// type p37b02ea1a.DemoUnit in package:github.com/starter-go/mimetypes/src/test/golang/unit
//
// id:com-37b02ea1af0e6fe2-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p37b02ea1af_unit_DemoUnit struct {
}

func (inst* p37b02ea1af_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-37b02ea1af0e6fe2-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p37b02ea1af_unit_DemoUnit) new() any {
    return &p37b02ea1a.DemoUnit{}
}

func (inst* p37b02ea1af_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p37b02ea1a.DemoUnit)
	nop(ie, com)

	


    return nil
}



// type p37b02ea1a.FindTypeUnit in package:github.com/starter-go/mimetypes/src/test/golang/unit
//
// id:com-37b02ea1af0e6fe2-unit-FindTypeUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p37b02ea1af_unit_FindTypeUnit struct {
}

func (inst* p37b02ea1af_unit_FindTypeUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-37b02ea1af0e6fe2-unit-FindTypeUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p37b02ea1af_unit_FindTypeUnit) new() any {
    return &p37b02ea1a.FindTypeUnit{}
}

func (inst* p37b02ea1af_unit_FindTypeUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p37b02ea1a.FindTypeUnit)
	nop(ie, com)

	
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*p37b02ea1af_unit_FindTypeUnit) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}



// type p37b02ea1a.ListTypesUnit in package:github.com/starter-go/mimetypes/src/test/golang/unit
//
// id:com-37b02ea1af0e6fe2-unit-ListTypesUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p37b02ea1af_unit_ListTypesUnit struct {
}

func (inst* p37b02ea1af_unit_ListTypesUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-37b02ea1af0e6fe2-unit-ListTypesUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p37b02ea1af_unit_ListTypesUnit) new() any {
    return &p37b02ea1a.ListTypesUnit{}
}

func (inst* p37b02ea1af_unit_ListTypesUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p37b02ea1a.ListTypesUnit)
	nop(ie, com)

	
    com.TM = inst.getTM(ie)


    return nil
}


func (inst*p37b02ea1af_unit_ListTypesUnit) getTM(ie application.InjectionExt)p85a4d026d.Manager{
    return ie.GetComponent("#alias-85a4d026daf77828ef49edb2adfd695e-Manager").(p85a4d026d.Manager)
}



// type p277ab6ae7.DemoTypes in package:github.com/starter-go/mimetypes/src/test/golang/unit/demotypes
//
// id:com-277ab6ae7e2e0e07-demotypes-DemoTypes
// class:class-85a4d026daf77828ef49edb2adfd695e-Registry
// alias:
// scope:singleton
//
type p277ab6ae7e_demotypes_DemoTypes struct {
}

func (inst* p277ab6ae7e_demotypes_DemoTypes) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-277ab6ae7e2e0e07-demotypes-DemoTypes"
	r.Classes = "class-85a4d026daf77828ef49edb2adfd695e-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p277ab6ae7e_demotypes_DemoTypes) new() any {
    return &p277ab6ae7.DemoTypes{}
}

func (inst* p277ab6ae7e_demotypes_DemoTypes) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p277ab6ae7.DemoTypes)
	nop(ie, com)

	


    return nil
}


