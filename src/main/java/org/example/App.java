package org.example;

import com.sun.jna.Memory;
import com.sun.jna.Native;
import org.json.JSONObject;

/**
 * Hello world!
 */
public class App {

    private static String LIB_PATH = App.class.getResource("/gojnalib/libjnawrapper.so").getPath();


    static GoJNALib INSTANCE = Native.loadLibrary(LIB_PATH, GoJNALib.class);

    public static void main(String[] args) {
        System.out.println(LIB_PATH);
        System.out.println(INSTANCE);
        System.out.println("Hello World!");

        JSONObject jsonObject1 = new JSONObject();
        jsonObject1.put("name", "John");
        jsonObject1.put("age", 30);
        jsonObject1.put("married", true);

        JSONObject jsonObject2 = new JSONObject();
        jsonObject2.put("name", "Jane");
        jsonObject2.put("age", 33);
        jsonObject2.put("married", false);
        String jsonString1 = jsonObject1.toString();

        System.out.println(jsonString1);
        String[] jsonStrings = {jsonString1, jsonObject2.toString()};
        for (String j : jsonStrings) {
            System.out.println(j);
        }

        // Convert Java strings to C strings and keep references to prevent GC
        Memory pointersMem = new Memory(Native.POINTER_SIZE * jsonStrings.length);
        //  Pointer[] pointers = new Pointer[jsonStrings.length];
        for (int i = 0; i < jsonStrings.length; i++) {
            byte[] bytes = (jsonStrings[i] + "\0").getBytes(); // C strings are null-terminated
            Memory stringMem = new Memory(bytes.length);
            stringMem.write(0, bytes, 0, bytes.length);
            pointersMem.setPointer(Native.POINTER_SIZE * i, stringMem);
        }


        INSTANCE.JNAWrapper(pointersMem, jsonStrings.length, jsonString1);
    }
}
