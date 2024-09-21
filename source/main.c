// Include the most common headers from the C standard library
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

// Include the main libnx system header, for Switch development
#include <switch.h>

#include "hello.h"

int main(int argc, char* argv[])
{
    consoleInit(NULL);
    padConfigureInput(1, HidNpadStyleSet_NpadStandard);
    PadState pad;
    padInitializeDefault(&pad);

    printf("Hello World!\n");

    // --------------------
    // If this is uncommented, padInitializeDefault() above crashes in hidInitializeNpad()
    HelloWorld();
    // --------------------

    while (appletMainLoop())
    {
        padUpdate(&pad);
        u64 kDown = padGetButtonsDown(&pad);

        if (kDown & HidNpadButton_Plus)
            break; // break in order to return to hbmenu
        consoleUpdate(NULL);
    }
    consoleExit(NULL);
    return 0;
}
