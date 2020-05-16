from ctypes import *

lib = cdll.LoadLibrary("./chudinov.so")

lib.IsSimple.argtypes = [c_int]
lib.IsSimple.restype = c_bool

lib.NextSimple.argtypes = [c_int]
lib.NextSimple.restype = c_int

simpleInt = c_int(11)
print("For integer 11")
print("Is simple: ")
print(lib.IsSimple(simpleInt))
print("Next simple: ")
print(lib.NextSimple(simpleInt))

simpleInt = c_int(14)
print("For integer 14")
print("Is simple: ")
print(lib.IsSimple(simpleInt))
print("Next simple: ")
print(lib.NextSimple(simpleInt))

