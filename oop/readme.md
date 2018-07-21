# 面向对象编程-继承
    
    type Goods struct{
        Name string
        Price int
    }
    
    type Book struct{
        Goods
        Writer string
    }
    
