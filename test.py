from ctypes import *

lib = cdll.LoadLibrary("./chudinov.so")

lib.IsSimple.argtypes = [c_int]
lib.IsSimple.restype = c_bool

simpleInt = c_int(2)
print(lib.IsSimple(simpleInt))

simpleInt = c_int(4)
print(lib.IsSimple(simpleInt))

lib.NextSimple.argtypes = [c_int]
lib.NextSimple.restype = c_int

simpleInt = c_int(2)
print(lib.NextSimple(simpleInt))

simpleInt = c_int(7)
print(lib.NextSimple(simpleInt))

