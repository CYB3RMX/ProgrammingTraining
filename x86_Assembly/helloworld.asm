section .data
    ; 0xa => \n
    message1 db "Displaying 9 stars", 0xa
    len equ $ - message1
    message2 times 9 db "*"

section .text
    global _start

_start:
    ; print("Displaying 9 stars")
    mov edx, len        ; message length
    mov ecx, message1   ; string to print
    mov ebx, 1          ; stdout
    mov eax, 4          ; sys_write
    int 0x80            ; system_call

    ; print("*********")
    mov edx, 9          ; length = 9
    mov ecx, message2   ; *********
    mov ebx, 1          ; stdout
    mov eax, 4          ; sys_write
    int 0x80            ; system_call

    ; exit()
    mov eax, 1 ; exit_call
    int 0x80