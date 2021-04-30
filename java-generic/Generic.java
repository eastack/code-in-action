package me.eastack;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class Generic {
    /**
     * 首先要清楚， extends 和 super 是用在范型类型“定义”中的
     */
    public void extendsGeneric() {
        // 定义范型类型 List<? extends Number>
        List<? extends Number> arr;
        // ArrayList<Integer> 是 List<? extends Number> 的子类
        arr = new ArrayList<Integer>();
        // ArrayList<Double> 是 List<? extends Number> 的子类
        arr = new ArrayList<Double>();
        // ArrayList<Float> 是 List<? extends Number> 的子类
        arr = new ArrayList<Float>();

        // set 报错，因为此时 arr0 有可能是
        // ArrayList<Float>
        // ArrayList<Float>
        // ArrayList<Float>
        // 三个子类型中的任意一个
        //arr0.set(0, Integer.valueOf(1));

        // get 获取到的值是 Number 类型，原因同上，取出类型为 Number 是安全的
        Number number = arr.get(0);
    }

    public void superGeneric() {
        // 定义范型类型 List<? super Number>
        List<? super Number> arr;
        // ArrayList<Object> 是 List<? super Number> 的子类
        arr = new ArrayList<Object>();
        // ArrayList<Serializable> 是 List<? super Number> 的子类
        arr = new ArrayList<Serializable>();

        // set 成功，因为此时 arr1 有可能是
        // ArrayList<Object>
        // ArrayList<Serializable>
        // 两个范型子类型中的任意一个
        arr.set(0, Integer.valueOf(1));
        // get 获取到 Object 类型，因为此时 arr1 有可能是
        // ArrayList<Object>
        // ArrayList<Serializable>
        // 两个范型子类型中的任意一个，
        // 而此处只有取出 Object 是安全的操作
        Object object = arr.get(0);
    }

    public void anyGeneric() {
        // 定义范型类型 List<? super Number>
        List<?> arr;
        // ArrayList<Object> 是 List<? super Number> 的子类
        arr = new ArrayList<Object>();
        // ArrayList<Serializable> 是 List<? super Number> 的子类
        arr = new ArrayList<Serializable>();
        // ArrayList<Integer> 是 List<? super Number> 的子类
        arr = new ArrayList<Integer>();

        // set 只有 null 成功其他均报错，因为此时 arr1 有可能是
        // ArrayList<Object>
        // ArrayList<Serializable>
        // ArrayList<Integer>
        // ArrayList<...>
        // ...
        // 等无数个范型子类型中的任意一个
        arr.set(0, null);
        // get 获取到 Object 类型，因为此时 arr1 有可能是
        // ArrayList<Object>
        // ArrayList<Serializable>
        // ArrayList<Integer>
        // ArrayList<...>
        // ...
        // 等无数个范型子类型中的任意一个，
        // 而此处只有取出 Object 是安全的操作
        Object o = arr.get(0);
    }

    public void oneGeneric() {
        // 定义范型类型 List<String>
        List<String> arr;
        // ArrayList<String> 是 List<String> 的子类，但范型类型确定且统一
        arr = new ArrayList<String>();
        arr = new LinkedList<String>();

        // set 成功，不多说
        arr.set(0, "generic");
        // get 成功，不多说
        String o = arr.get(0);
    }
}
