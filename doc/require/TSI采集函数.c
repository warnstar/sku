void *TSI_Func(void *arg) //返回值类型是void *
{
	
	int i = 0;
	float temp =0.0;
	char times = 4;
	FILE *fd;
  time_t now;
  time(&now);// 等同于now = time(NULL)
	
	sleep(2); //这行要优化掉 
	
	//刚开始读取的都是0，要丢弃
	while (temp< 2.0 ) {
		
	temp = TSI_getData();
	printf("等待不再为0...\n");
	sleep(1);
	
}	


	 while (temp < 500+10 || times-- ) 
	 {
		printf("等待达到500，此时%0.1f....%d\n",temp,times);
		temp = TSI_getData();
		sleep(1);

  }
  
	times = 4;
	
	while (temp > 500 || times-- ) 
	{
		printf("等待下降到500，此时%0.1f....%d\n",temp,times);
		temp = TSI_getData();
		sleep(1);
		
	}

	g_Step = 1;
	printf("开始采集\n");
  
  if (g_Switch) 
  {
  	fd=fopen("data_TSI_test.txt","a+");
  }
  else 
  {
  	fd=fopen("data_TSI.txt","a+");
  	
  }
	
	
  if(fd==NULL)//如果失败了
  {
      printf("打开 错误！");
      pthread_exit(NULL) ;//中止程序
  }

 	 	 
  	printf("已采集\n");
  	
  	//13
    while (g_Step && i < SampleNumMaxAll && temp > 13)  { //TSI 最多采样6000个数据,小于11的时候停止采样，是为了能够在step9中计算14-50
 		temp = TSI_getData();
 		while (temp < 1.0) {
 			sleep(1);
 			temp = TSI_getData();
 			
 		} 
		time(&now);
		fprintf(fd,"%u:%f\n",(unsigned int)now,temp);
		sleep(1);
		i++;
		if (!(i%10)) {printf("%d\n",i);  fflush(fd);}
    }
    
    printf("采集完成,总数量%d\n",i);
    g_Step = 0; //主要TSI停止，则线程也被强制停止
    fflush(fd);
    fclose(fd);
    TSI_stop();
    TSI_getData();
		return((void *)0); 

}