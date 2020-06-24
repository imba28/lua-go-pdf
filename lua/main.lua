local ffi = require("ffi")
local pdf = ffi.load('./pdf.so')
local io = require "io"

ffi.cdef([[
typedef long long GoInt64;
typedef GoInt64 GoInt;

typedef struct { void *t; void *v; } GoInterface;
typedef struct { const char *p; GoInt n; } GoString;

/* Return type for render */
struct render_return {
	char *r0;
	GoInt r1;
};

extern struct render_return render(GoString p0);
]]);

local typeString = ffi.metatype("GoString", {})
local templateName = "template.html"
local templateGoString = typeString(templateName, #templateName)

local result = pdf.render(templateGoString)
local content = ffi.string(result.r0, result.r1)

file = io.open("invoice.pdf", "w")
io.output(file)
io.write(content)
io.close(file)