/*
 * numlockbios.c - Check the NumLock status from the BIOS data area
 *
 * Borrowed from the kbd package of SuSE Linux
 * Copyright (C) 1995-2025 SuSE GmbH
 * Copyright (C) 2024      YggdrasilSoft, LLC
 * 
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 * 
 * CHANGES
 * 2025-12-29 - Adapted for MidgardOS and removed all goto statements.
 */

#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

int finish(int fdmem, int exit_code) {
    close(fdmem);
    printf("unknown\n");
    exit(exit_code);
}

int main() {
    #define BIOS_DATA_AREA  0x400
    #define BDA_KEYBOARD_STATUS_FLAGS_4 0x97
    #define BDA_KSF4_NUMLOCK_MASK 0x02

    int fdmem;
    char c;
    errno=0;

    fdmem = open("/dev/mem", O_RDONLY);

    if (fdmem < 0) {
        fprintf(stderr, "Couldn't open /dev/mem; %s\n", strerror(errno));
        finish(fdmem, errno);
    }

    if (lseek(fdmem, BIOS_DATA_AREA + BDA_KEYBOARD_STATUS_FLAGS_4, SEEK_SET) == (off_t) -1) {
        fprintf(stderr, "Failed to seek /dev/mem: %s\n", strerror(errno));
        finish(fdmem, errno);
    }

    if (read (fdmem, &c, sizeof(char)) == -1) {
        fprintf(stderr, "Failed to read /dev/mem: %s\n", strerror(errno));
        finish(fdmem, errno);
    }

    if (c & BDA_KSF4_NUMLOCK_MASK) {
        printf("on\n");
    } else {
        printf("off\n");
    }

    finish(fdmem, 0);
    return 0;
}

/* vim: set ts=4 sw=4 noet: */
