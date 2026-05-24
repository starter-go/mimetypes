package main4mimetypes
import (
    p52d293b71 "github.com/starter-go/i18n"
    p85a4d026d "github.com/starter-go/mimetypes"
    p92665872b "github.com/starter-go/mimetypes/src/main/golang/lib"
     "github.com/starter-go/application"
)

// type p92665872b.DefaultTypeManager in package:github.com/starter-go/mimetypes/src/main/golang/lib
//
// id:com-92665872bf497637-lib-DefaultTypeManager
// class:
// alias:alias-85a4d026daf77828ef49edb2adfd695e-Manager
// scope:singleton
//
type p92665872bf_lib_DefaultTypeManager struct {
}

func (inst* p92665872bf_lib_DefaultTypeManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-92665872bf497637-lib-DefaultTypeManager"
	r.Classes = ""
	r.Aliases = "alias-85a4d026daf77828ef49edb2adfd695e-Manager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p92665872bf_lib_DefaultTypeManager) new() any {
    return &p92665872b.DefaultTypeManager{}
}

func (inst* p92665872bf_lib_DefaultTypeManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p92665872b.DefaultTypeManager)
	nop(ie, com)

	
    com.Sources = inst.getSources(ie)
    com.I18n = inst.getI18n(ie)


    return nil
}


func (inst*p92665872bf_lib_DefaultTypeManager) getSources(ie application.InjectionExt)[]p85a4d026d.Registry{
    dst := make([]p85a4d026d.Registry, 0)
    src := ie.ListComponents(".class-85a4d026daf77828ef49edb2adfd695e-Registry")
    for _, item1 := range src {
        item2 := item1.(p85a4d026d.Registry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p92665872bf_lib_DefaultTypeManager) getI18n(ie application.InjectionExt)p52d293b71.Service{
    return ie.GetComponent("#alias-52d293b7197ae4adfab092ade2e718a2-Service").(p52d293b71.Service)
}


