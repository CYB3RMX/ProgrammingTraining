section .data
    user_msg db 'Enter a number: '
    len_user_msg equ $ - user_msg
    disp_msg db 'You have entered: '
    len_disp_msg equ $ - disp_msg

section .bss
    num resb 5

section .text
    global _start

_start:
    ; print("Enter a number: ")
    mov edx, len_user_msg
    mov ecx, user_msg
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; num = input()
    mov edx, 5      ; 4 bytes int | 1 byte info
    mov ecx, num
    mov ebx, 2      ; stdin
    mov eax, 3      ; sys_read
    int 0x80

    ; print("You have entered: {num}")
    mov edx, len_disp_msg
    mov ecx, disp_msg
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; print(num)
    mov edx, 5
    mov ecx, num,
    mov ebx, 1
    mov eax, 4
    int 0x80

    ; exit()
    mov eax, 1
    int 0x80