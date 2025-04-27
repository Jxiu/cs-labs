package com.github.jxiu.flink.stock;

/**
 * 描述： 股票信息
 *
 * <pre>
 * HISTORY
 * ********************
 *  ID   DATE
 *  1    2024-12-16
 * ********************
 * </pre>
 *
 * @author jxiu
 * @version 1.0
 */
public class Stock {

    private String symbol;
    private String name;
    private double price;
    private long ts;

    public Stock() {
        // default constructor
    }

    public Stock(String symbol, String name, double price) {
        this.symbol = symbol;
        this.name = name;
        this.price = price;
    }

    public String getSymbol() {
        return symbol;
    }

    public void setSymbol(String symbol) {
        this.symbol = symbol;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }
}
