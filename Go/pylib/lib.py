from ctypes import cdll , c_char_p
dll = cdll.LoadLibrary("./lib")

dll.retutf8.restype = c_char_p
dll.retascii.restype = c_char_p

print(dll.retutf8())
print(dll.retascii())


#dll.retutf8(c_char_p("Merhaba dünya")) # bu eleman ile go'ya str atabiliriz