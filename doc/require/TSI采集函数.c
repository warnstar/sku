void *TSI_Func(void *arg) //����ֵ������void *
{
	
	int i = 0;
	float temp =0.0;
	char times = 4;
	FILE *fd;
  time_t now;
  time(&now);// ��ͬ��now = time(NULL)
	
	sleep(2); //����Ҫ�Ż��� 
	
	//�տ�ʼ��ȡ�Ķ���0��Ҫ����
	while (temp< 2.0 ) {
		
	temp = TSI_getData();
	printf("�ȴ�����Ϊ0...\n");
	sleep(1);
	
}	


	 while (temp < 500+10 || times-- ) 
	 {
		printf("�ȴ��ﵽ500����ʱ%0.1f....%d\n",temp,times);
		temp = TSI_getData();
		sleep(1);

  }
  
	times = 4;
	
	while (temp > 500 || times-- ) 
	{
		printf("�ȴ��½���500����ʱ%0.1f....%d\n",temp,times);
		temp = TSI_getData();
		sleep(1);
		
	}

	g_Step = 1;
	printf("��ʼ�ɼ�\n");
  
  if (g_Switch) 
  {
  	fd=fopen("data_TSI_test.txt","a+");
  }
  else 
  {
  	fd=fopen("data_TSI.txt","a+");
  	
  }
	
	
  if(fd==NULL)//���ʧ����
  {
      printf("�� ����");
      pthread_exit(NULL) ;//��ֹ����
  }

 	 	 
  	printf("�Ѳɼ�\n");
  	
  	//13
    while (g_Step && i < SampleNumMaxAll && temp > 13)  { //TSI ������6000������,С��11��ʱ��ֹͣ��������Ϊ���ܹ���step9�м���14-50
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
    
    printf("�ɼ����,������%d\n",i);
    g_Step = 0; //��ҪTSIֹͣ�����߳�Ҳ��ǿ��ֹͣ
    fflush(fd);
    fclose(fd);
    TSI_stop();
    TSI_getData();
		return((void *)0); 

}