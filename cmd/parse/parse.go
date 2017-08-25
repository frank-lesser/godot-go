package main

/*
This is meant to parse the godot documentation to generate Go structs of all the godot
objects.
*/

import (
	"encoding/xml"
	"fmt"
	"github.com/shadowapex/godot-go/templates"
	"io/ioutil"
	"net/http"
)

const classesUrl = "https://raw.githubusercontent.com/godotengine/godot/master/doc/base/classes.xml"

func main() {

	// Fetch the classes XML
	resp, err := http.Get(classesUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the XML into a defined structure. This structure is generated by
	// cmd/generate.sh.
	var classes templates.GDDoc
	xml.Unmarshal(body, &classes)

	// Loop through all classes
	for _, class := range classes.GDClass {
		fmt.Println("Name:", class.AttrName)
		fmt.Println("  Category:", class.AttrCategory)
		fmt.Println("  Inherits:", class.AttrInherits)
		fmt.Println("  Description:", class.GDDescription.Text)
		fmt.Println("  Constants:")
		if class.GDConstants != nil {
			for _, constant := range class.GDConstants.GDConstant {
				fmt.Println("    Constant:", constant)
			}
		}
		fmt.Println("  Members:")
		if class.GDMembers != nil {
			for _, member := range class.GDMembers.GDMember {
				fmt.Println("    Member:", member)
			}
		}
		fmt.Println("  Methods:")
		if class.GDMethods != nil {
			for _, method := range class.GDMethods.GDMethod {
				fmt.Println("    Method:")
				fmt.Println("      Name:", method.AttrName)
				fmt.Println("      Description:", method.GDDescription.Text)
				fmt.Println("      Qualifiers:", method.AttrQualifiers)
				fmt.Println("      Arguments:")
				for _, arg := range method.GDArgument {
					fmt.Println("        Name:", arg.AttrName)
					fmt.Println("        Default:", arg.AttrDefault)
					fmt.Println("        Index:", arg.AttrIndex)
					fmt.Println("        Type:", arg.AttrType)
				}
				fmt.Println("      Returns:")
				if method.GDReturn != nil {
					fmt.Println("        Name:", method.GDReturn.AttrType)
				}
			}
		}
		fmt.Println("  Signals:")
		if class.GDSignals != nil {
			for _, signal := range class.GDSignals.GDSignal {
				fmt.Println("    Signal:", signal)
			}
		}
		fmt.Println("  Theme Items:")
		if class.GDTheme_items != nil {
			for _, theme := range class.GDTheme_items.GDTheme_item {
				fmt.Println("    Theme:", theme)
			}
		}
	}
}