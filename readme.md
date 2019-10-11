# Jenkinsfile文件格式化工具

对简单的Jenkinsfile文件可以格式化

## 对满足如下条件的Jenkinsfile文件可以格式化
1. 每一行最多有一个花括号，如果有花括号，则以花括号结尾（结尾允许有行注释或者空白字符）  
例如下面这个可以格式化：
```
pipeline { //行末可以有行注释或空白字符
stages {
stage('stage1') {
}
}
}
```
而这个没法格式化：
```
pipeline {stages {stage('stage1') {}}}
```
2. 块注释不会被格式化，但是如果有块注释，则要求块注释开头"/*"出现在每行的最左边，块注释的结束"*/"出现在另一行的最右边（如果你想块注释出现在同一行，直接用行注释"//"吧），否则可能格式化失败  

例如允许下面的块注释：
```
pipeline {
    stages {
        /*stage('stage1') {
        }*/
    }
}
```
不允许下面2种的块注释：
```
pipeline {
    stages {/*
        stage('stage1') {
        }*/
    }
}
```
```
pipeline {
    stages {
        /*stage('stage1') {}*/
    }
}
```
