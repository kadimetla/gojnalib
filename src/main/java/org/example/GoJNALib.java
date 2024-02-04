package org.example;

import com.sun.jna.Library;
import com.sun.jna.Memory;

public interface GoJNALib extends Library {
    void JNAWrapper(Memory strings, int length, String apiRawInput);
}
