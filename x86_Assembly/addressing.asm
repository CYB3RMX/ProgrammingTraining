section .data
    message1 db 'Enter a number: '
    len_message1 equ $ - message1
    message2 db 'Before addressing: '
    len_message2 equ $ - message2
    message3 db 'After addressing: '
    len_message3 equ $ - message3

section .bss
    number resb 5

section .text
    global _start

_start:
    ; print("Enter a number: ")
    mov edx, len_message1
    mov ecx, message1,
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; number = input()
    mov edx, 5
    mov ecx, number,
    mov ebx, 2
    mov eax, 3
    int 0x80

    ; print("Before addressing: ")
    mov edx, len_message2
    mov ecx, message2
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; print(number)
    mov edx, 5
    mov ecx, number
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; addressing => number = 5555
    mov [number], dword '5555'

    ; print("After addressing: ")
    mov edx, len_message3
    mov ecx, message3
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; print(number)
    mov edx, 5
    mov ecx, number
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; exit()
    mov eax, 1
    int 0x80