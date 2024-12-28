## 使用Jmap可以做到吗？

可以使用 jmap -histo:live [pid] 获取当前堆中存活对象数量以及空间占用情况，每次调用都需要扫描整个堆，对程序性能影响较大。

## 什么是JVMTI？

**JVMTI**（Java Virtual Machine Tool Interface）是一个用于与 Java 虚拟机（JVM）进行交互的接口，允许开发者或工具在运行时监控、调试、分析和操作 JVM 内部的各种状态。JVMTI 提供了底层的功能，可以实现对虚拟机的细粒度控制，通常用于性能分析、垃圾回收调优、堆栈跟踪、线程管理等任务。

### 主要特点：

1. **事件通知**：JVMTI 支持在 JVM 中触发特定事件（如线程启动、堆内存分配、垃圾回收等）并通过回调通知开发者。
2. **线程和内存分析**：提供对 JVM 内部线程和内存状态的访问，可以获取堆栈、堆信息、线程状态等。
3. **代码和方法监控**：可以监控方法的调用、返回，插桩字节码，获取堆栈信息等。
4. **堆栈跟踪**：支持堆栈跟踪功能，可以用于调试和分析。
5. **无缝调试**：可以在不改变原始代码的情况下插入监控逻辑，帮助开发人员进行深度的 JVM 调试。
6. 

```c
typedef struct {
                              /*   50 : VM Initialization Event */
    jvmtiEventVMInit VMInit;
                              /*   51 : VM Death Event */
    jvmtiEventVMDeath VMDeath;
                              /*   52 : Thread Start */
    jvmtiEventThreadStart ThreadStart;
                              /*   53 : Thread End */
    jvmtiEventThreadEnd ThreadEnd;
                              /*   54 : Class File Load Hook */
    jvmtiEventClassFileLoadHook ClassFileLoadHook;
                              /*   55 : Class Load */
    jvmtiEventClassLoad ClassLoad;
                              /*   56 : Class Prepare */
    jvmtiEventClassPrepare ClassPrepare;
                              /*   57 : VM Start Event */
    jvmtiEventVMStart VMStart;
                              /*   58 : Exception */
    jvmtiEventException Exception;
                              /*   59 : Exception Catch */
    jvmtiEventExceptionCatch ExceptionCatch;
                              /*   60 : Single Step */
    jvmtiEventSingleStep SingleStep;
                              /*   61 : Frame Pop */
    jvmtiEventFramePop FramePop;
                              /*   62 : Breakpoint */
    jvmtiEventBreakpoint Breakpoint;
                              /*   63 : Field Access */
    jvmtiEventFieldAccess FieldAccess;
                              /*   64 : Field Modification */
    jvmtiEventFieldModification FieldModification;
                              /*   65 : Method Entry */
    jvmtiEventMethodEntry MethodEntry;
                              /*   66 : Method Exit */
    jvmtiEventMethodExit MethodExit;
                              /*   67 : Native Method Bind */
    jvmtiEventNativeMethodBind NativeMethodBind;
                              /*   68 : Compiled Method Load */
    jvmtiEventCompiledMethodLoad CompiledMethodLoad;
                              /*   69 : Compiled Method Unload */
    jvmtiEventCompiledMethodUnload CompiledMethodUnload;
                              /*   70 : Dynamic Code Generated */
    jvmtiEventDynamicCodeGenerated DynamicCodeGenerated;
                              /*   71 : Data Dump Request */
    jvmtiEventDataDumpRequest DataDumpRequest;
                              /*   72 */
    jvmtiEventReserved reserved72;
                              /*   73 : Monitor Wait */
    jvmtiEventMonitorWait MonitorWait;
                              /*   74 : Monitor Waited */
    jvmtiEventMonitorWaited MonitorWaited;
                              /*   75 : Monitor Contended Enter */
    jvmtiEventMonitorContendedEnter MonitorContendedEnter;
                              /*   76 : Monitor Contended Entered */
    jvmtiEventMonitorContendedEntered MonitorContendedEntered;
                              /*   77 */
    jvmtiEventReserved reserved77;
                              /*   78 */
    jvmtiEventReserved reserved78;
                              /*   79 */
    jvmtiEventReserved reserved79;
                              /*   80 : Resource Exhausted */
    jvmtiEventResourceExhausted ResourceExhausted;
                              /*   81 : Garbage Collection Start */
    jvmtiEventGarbageCollectionStart GarbageCollectionStart;
                              /*   82 : Garbage Collection Finish */
    jvmtiEventGarbageCollectionFinish GarbageCollectionFinish;
                              /*   83 : Object Free */
    jvmtiEventObjectFree ObjectFree;
                              /*   84 : VM Object Allocation */
    jvmtiEventVMObjectAlloc VMObjectAlloc;
} jvmtiEventCallbacks;
```

其中我们先关注
jvmtiEventObjectFree ObjectFree 对象释放
jvmtiEventVMObjectAlloc VMObjectAlloc 虚拟机对象分配


```c
typedef void (JNICALL *jvmtiEventObjectFree)

    (jvmtiEnv *jvmti_env,

     jlong tag);
```

- **`jvmtiEnv *jvmti_env`**
    
    - 这是 `JVMTI` 环境指针，它是你访问 JVMTI 功能的入口。所有的 `JVMTI` API 调用都需要这个指针来进行操作。
    - tag 表示被释放对象在JVM中的内存位置。


```c
typedef void (JNICALL *jvmtiEventVMObjectAlloc) (jvmtiEnv *jvmti_env, JNIEnv* jni_env, jthread thread, jobject object, jclass object_klass, jlong size);
```


- **`jvmtiEnv *jvmti_env`**
    
    - 这是 `JVMTI` 环境指针，它是你访问 JVMTI 功能的入口。所有的 `JVMTI` API 调用都需要这个指针来进行操作。
- **`JNIEnv* jni_env`**
    
    - 这是 `JNI` 环境指针，提供访问 Java 本地方法接口的功能。通过它，你可以调用 Java 方法、访问 Java 对象等。通常在 JNI 或 JVMTI 回调中都需要这个指针来与 Java 环境交互。
- **`jthread thread`**
    
    - 这是触发对象分配的线程。在 `VMObjectAlloc` 事件中，你会知道哪个线程分配了这个对象。它是一个指向线程对象的引用，可以通过 JNI 来访问该线程的相关信息。
- **`jobject object`**
    
    - 这是指向新分配对象的引用。这个 `jobject` 是一个 Java 对象，它是新创建的对象的实例。你可以通过 `JNIEnv` 来访问对象的字段或调用该对象的方法。
        
    - **`jobject` 的字段**：在 Java 中，`jobject` 代表一个对象的引用，它本身并没有字段，因为它只是一个引用类型。你不能直接访问 `jobject` 中的字段。要访问对象的字段，你需要使用 JNI 提供的函数，比如 `GetObjectField`，`GetIntField` 等。这些函数可以根据对象的类和字段名/ID 来提取字段的值。
        
    - **注意**：`jobject` 是一个 Java 对象的引用，具体是哪个类的对象，可以通过 `object_klass` 来确认。如果你需要进一步操作该对象的字段（例如访问其成员变量），你可以在回调函数中使用 `JNIEnv` 来执行操作。
        
- **`jclass object_klass`**
    
    - 这是对象所属的类的引用。在 Java 中，每个对象都有一个类，`object_klass` 提供了该对象的类信息。你可以使用它来获取对象的类名、父类、方法、字段等信息。通常可以通过 `GetClassName` 等 JNI 函数来获取类的相关信息。
- **`jlong size`**
    
    - 这是新分配对象的大小，单位是字节。它表示对象在堆内存中分配的内存大小。这个信息对于内存监控、性能分析等场景非常有用。



> [!NOTE] jvmtiEventVMObjectAlloc发生的时机
> 当 JVM 为一个新对象分配内存时，首先会分配足够的内存空间，然后在这块内存中执行构造函数，初始化对象的字段等操作。`jvmtiEventVMObjectAlloc` 事件恰好在内存分配完成、构造函数执行之前触发，因此你在这个回调中会获得对象的引用、大小以及类信息，但还不能访问构造函数中初始化的字段或对象的实际内容。


注册对象分配监听事件示例代码

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <jvmti.h>

// 错误处理函数
void check_jvmti_error(jvmtiError err, const char *msg) {
    if (err != JVMTI_ERROR_NONE) {
        fprintf(stderr, "JVMTI error: %d (%s)\n", err, msg);
        exit(1);
    }
}

// 回调函数：对象分配时触发
void JNICALL object_alloc_callback(jvmtiEnv *jvmti_env,
                                    JNIEnv *jni_env,
                                    jthread thread,
                                    jobject object,
                                    jclass object_klass,
                                    jlong size) {
    // 获取对象类的签名
    char *class_signature = NULL;
    char *generic_signature = NULL;

    jvmtiError err = (*jvmti_env)->GetClassSignature(jvmti_env, object_klass, &class_signature, &generic_signature);
    check_jvmti_error(err, "GetClassSignature");

    // 打印对象的类签名和大小
    printf("Object of class %s allocated with size %lld bytes.\n", class_signature, size);

    // 释放获取的签名字符串
    (*jvmti_env)->Deallocate(jvmti_env, (unsigned char *)class_signature);
    (*jvmti_env)->Deallocate(jvmti_env, (unsigned char *)generic_signature);
}

// 设置 JVMTI 回调函数
void setup_jvmti_callbacks(jvmtiEnv *jvmti_env) {
    jvmtiEventCallbacks callbacks;
    memset(&callbacks, 0, sizeof(callbacks));

    // 设置对象分配事件的回调函数
    callbacks.VMObjectAlloc = object_alloc_callback;

    // 注册回调函数
    jvmtiError err = (*jvmti_env)->SetEventCallbacks(jvmti_env, &callbacks, sizeof(callbacks));
    check_jvmti_error(err, "SetEventCallbacks");

    // 在启用对象分配事件通知前，需要先检测是否具有对应能力，否则会报错 JVMTI_ERROR_MUST_POSSESS_CAPABILITY
    jvmtiCapabilities capabilities;
    memset(&capabilities, 0, sizeof(capabilities));

    err = (*jvmti_env)->GetCapabilities(jvmti_env, &capabilities);
    check_jvmti_error(err, "GetCapabilities");

    if (!(capabilities.can_generate_vm_object_alloc_events)) {
        capabilities.can_generate_vm_object_alloc_events = 1;
        err = (*jvmti_env)->AddCapabilities(jvmti_env, &capabilities);
        check_jvmti_error(err, "AddCapabilities");
    }

    // 启用对象分配事件通知
    err = (*jvmti_env)->SetEventNotificationMode(jvmti_env, JVMTI_ENABLE, JVMTI_EVENT_VM_OBJECT_ALLOC, NULL);
    check_jvmti_error(err, "SetEventNotificationMode");
}

// Agent 加载函数：JVM 启动时调用
JNIEXPORT jint JNICALL Agent_OnLoad(JavaVM *vm, char *options, void *reserved) {
    jvmtiEnv *jvmti_env = NULL;
    JNIEnv *jni_env = NULL;
    jvmtiError err;

    // 获取 JVMTI 环境
    err = (*vm)->GetEnv(vm, (void **)&jvmti_env, JVMTI_VERSION_1_0);
    if (err != JVMTI_ERROR_NONE) {
        fprintf(stderr, "Unable to get JVMTI environment: %d\n", err);
        return JNI_ERR;  // 返回错误代码
    }

    // 设置 JVMTI 事件回调
    setup_jvmti_callbacks(jvmti_env);

    // 打印调试信息，确认 agent 已加载
    printf("Agent loaded successfully!\n");

    return JNI_OK;  // 返回成功代码
}

// Agent 卸载函数：JVM 停止时调用
JNIEXPORT void JNICALL Agent_OnUnload(JavaVM *vm) {
    // 如果需要在 JVM 卸载时执行清理工作，可以在这里进行
    printf("Agent unloading...\n");
}
```



```sh
# 编译windows dll
gcc -shared -I"D:\xxx\jdk8\include" -I"D:\xxx\jdk8\include\win32" -o monitor.dll monitor.c -L"D:\xxx\jdk8\lib" -ljvm

# 编译linux so
gcc -shared -I"D:\xxx\jdk8\include" -I"D:\xxx\jdk8\include\win32" -o monitor.so monitor.c -L"D:\xxx\jdk8\lib" -ljvm

# 启动JVM
java.exe  -agentpath:"D:/xxx/JavaObjetMonitor/monitor.dll"  -jar D:\xxx\Demo6-1.0-SNAPSHOT.jar
```



但是发现并没有监听到业务对象的创建，监听到的都是一些java.lang下面的对象。
```
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/String; allocated with size 24 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class Ljava/lang/StackTraceElement; allocated with size 32 bytes.
Object of class [Ljava/lang/StackTraceElement; allocated with size 120 bytes.
Object of class Ljava/lang/invoke/MemberName; allocated with size 56 bytes.
Object of class [Ljava/lang/Class; allocated with size 24 bytes.
```



通过[VMObjectAlloc文档](https://docs.oracle.com/javase/8/docs/platform/jvmti/jvmti.html#VMObjectAlloc)对其的介绍我可以可以得知VMObjectAlloc只会在针对那些没有字节码的对象内存分配的时候才会触发，对于存在字节码的对象（我们Java代码中定义的对象）我们应该使用字节码插装实现对象分配监控，相应的JNI对象的检测应该使用JNI function interception机制（我们这里不探究）。

> Sent when a method causes the virtual machine to allocate an Object visible to Java programming language code and the allocation is not detectable by other intrumentation mechanisms. Generally object allocation should be detected by instrumenting the bytecodes of allocating methods. Object allocation generated in native code by JNI function calls should be detected using [JNI function interception](https://docs.oracle.com/javase/8/docs/platform/jvmti/jvmti.html#jniIntercept). Some methods might not have associated bytecodes and are not native methods, they instead are executed directly by the VM. These methods should send this event. Virtual machines which are incapable of bytecode instrumentation for some or all of their methods can send this event.


## 如何通过字节码插装实现对象分配监控？

### 什么是javaagent？

`javaagent` 是通过在启动 JVM 时传递 `-javaagent` 参数来实现的，允许我们在程序启动时执行一些额外的代码，通常用于对类进行插装（字节码修改）。

当指定 `-javaagent` 参数时，JVM 会启动一个代理（agent），并为它提供一些功能，如：修改字节码、修改类加载过程、监控方法执行等。

#### 2. **运行流程**

`javaagent` 的运行流程主要包括以下几个步骤：

#### 2.1 **定义 Agent**

一个 `javaagent` 通常是一个包含 `premain` 方法的 Java 类。这个方法会在 JVM 启动时自动调用。

java

复制代码

`import java.lang.instrument.Instrumentation;  public class MyAgent {     public static void premain(String agentArgs, Instrumentation inst) {         System.out.println("Agent is running!");         // 可以注册 ClassFileTransformer 来修改字节码     } }`

#### 2.2 **启动时加载 Agent**

通过在启动 JVM 时使用 `-javaagent` 参数指定该 agent：

sh

复制代码

`java -javaagent:/path/to/agent.jar -jar myApplication.jar`

JVM 启动时会查找并加载 `agent.jar` 文件，然后调用其中的 `premain` 方法。

#### 2.3 **premain 方法**

`premain` 方法类似于 `main` 方法，但是它是在应用程序的主 `main` 方法之前被调用的。`premain` 方法有两个参数：

- `agentArgs`：从命令行传入的参数（如果有的话）。
- `inst`：一个 `Instrumentation` 对象，通过它可以对类进行修改或监控。

#### 2.4 **注册 ClassFileTransformer**

`Instrumentation` 对象可以用来注册一个 `ClassFileTransformer`，它可以拦截所有加载的类，并修改它们的字节码。

```java
inst.addTransformer(new ClassFileTransformer() {
    @Override
    public byte[] transform(ClassLoader loader, String className, Class<?> classBeingRedefined, ProtectionDomain protectionDomain, byte[] classfileBuffer) {
        // 这里可以修改类的字节码
        return classfileBuffer; // 返回修改后的字节码
    }
});
```

这个方法会在每个类加载时调用，你可以在其中对类字节码进行修改或插装。


### 如何利用javaagent监控对象创建

下面是一个使用Javassist(Javassist 是一个开源的 Java 字节码操作库，允许在运行时或编译时动态地修改、生成和分析 Java 类的字节码。它通过简化字节码操作，使开发者能够轻松地进行类的增强、代理和方法插桩等操作。)进行对象创建监控的例子

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>ObjectCreationAgent</artifactId>
    <version>1.0-SNAPSHOT</version>
    <packaging>jar</packaging>

    <dependencies>
        <!-- javassist for bytecode instrumentation -->
        <dependency>
            <groupId>org.javassist</groupId>
            <artifactId>javassist</artifactId>
            <version>3.28.0-GA</version>
        </dependency>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-jar-plugin</artifactId>
                <version>3.1.0</version>
                <configuration>
                    <archive>
                        <manifestEntries>
                            <Premain-Class>com.lin.ObjectCreationAgent</Premain-Class>  <!-- 指定你的代理类 -->
                            <Can-Redefine-Classes>true</Can-Redefine-Classes>  <!-- 可选的，允许重新定义类 -->
                        </manifestEntries>
                    </archive>
                </configuration>
            </plugin>
            <!--将依赖也打包进JAR包里-->
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-shade-plugin</artifactId>
                <version>3.2.1</version>
                <executions>
                    <execution>
                        <phase>package</phase>
                        <goals>
                            <goal>shade</goal>
                        </goals>
                        <configuration>
                            <createDependencyReducedPom>false</createDependencyReducedPom>
                            <shadedArtifactAttached>false</shadedArtifactAttached>
                        </configuration>
                    </execution>
                </executions>
            </plugin>
        </plugins>
    </build>
</project>

```


```java
package com.lin;  
  
import javassist.*;  
  
import java.lang.instrument.*;  
  
public class ObjectCreationAgent {  
    public static void premain(String agentArgs, Instrumentation inst) throws Exception {  
        System.out.println("ObjectCreationAgent.premain");  
        ClassPool pool = ClassPool.getDefault();  
        pool.appendClassPath(new LoaderClassPath(ObjectCreationAgent.class.getClassLoader()));  
  
        inst.addTransformer((loader, className, classBeingRedefined, protectionDomain, classfileBuffer) -> {  
            // 只插装自己类，随意插装其他类可能造成启动失败  
            if (!className.startsWith("com/lin")) {  
                return classfileBuffer;  
            }  
            try {  
                CtClass cc = pool.makeClass(new java.io.ByteArrayInputStream(classfileBuffer));  
                System.out.println("addTransformer " + cc.getName());  
                // 在构造方法中插入字节码  
                CtConstructor[] constructors = cc.getDeclaredConstructors();  
                for (CtConstructor constructor : constructors) {  
                    // 所有的构造函数前面都会去调用onObjectCreated方法  
                    constructor.insertBefore("{ com.lin.ObjectCreationAgent.onObjectCreated(\"" + cc.getName() + "\"); }");  
                }  
                return cc.toBytecode();  
            } catch (Throwable e) {  
                e.printStackTrace();  
            }  
            return classfileBuffer;  
        });  
    }  
  
    // 静态方法  
    public static void onObjectCreated(String className) {  
        System.out.println("Object created: " + className);  
    }  
}
```



java.exe  -javaagent:"D:\xxx\ObjectCreationAgent-1.0-SNAPSHOT.jar"  -jar D:\xxx\Demo6-1.0-SNAPSHOT.jar

![[Pasted image 20241225203410.png]]

==需要注意的是 javaagent和主程序同为JVM程序，共用堆栈和CPU，需要仔细编写，但是相对agentpath更为安全，后者可以直接接触JVM C++虚拟机运行代码，API更为底层==