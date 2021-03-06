package clang

// #include "./clang-c/CXCompilationDatabase.h"
// #include "go-clang.h"
import "C"

// CompileCommand represents the command line invocation to compile a specific file.
type CompileCommand struct {
	c C.CXCompileCommand
}

// Directory get the working directory where the CompileCommand was executed from
func (cc CompileCommand) Directory() string {
	o := cxstring{C.clang_CompileCommand_getDirectory(cc.c)}
	defer o.Dispose()

	return o.String()
}

// Filename get the filename associated with the CompileCommand.
func (cc CompileCommand) Filename() string {
	o := cxstring{C.clang_CompileCommand_getFilename(cc.c)}
	defer o.Dispose()

	return o.String()
}

// NumArgs get the number of arguments in the compiler invocation.
func (cc CompileCommand) NumArgs() uint32 {
	return uint32(C.clang_CompileCommand_getNumArgs(cc.c))
}

// Arg get the I'th argument value in the compiler invocations
//
// Invariant :
// - argument 0 is the compiler executable
func (cc CompileCommand) Arg(i uint32) string {
	o := cxstring{C.clang_CompileCommand_getArg(cc.c, C.uint(i))}
	defer o.Dispose()

	return o.String()
}
