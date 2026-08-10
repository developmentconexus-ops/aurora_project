package ports

import "context"

type ExportProtection interface { Protect(context.Context,[]byte,[]byte)([]byte,error); Unprotect(context.Context,[]byte,[]byte)([]byte,error) }
