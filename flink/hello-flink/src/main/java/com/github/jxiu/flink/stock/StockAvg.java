package com.github.jxiu.flink.stock;

import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;

/**
 * 描述：
 *
 * <pre>
 * HISTORY
 * *********************
 *  ID   DATE
 *  1    2024-12-16
 * *********************
 * </pre>
 *
 * @author jxiu
 * @version 1.0
 */
public class StockAvg {

    public static void main(String[] args) {
        StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();

        DataStream<String> stockStream = env.fromData(new Stock())
    }
}
