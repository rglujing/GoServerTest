#include<stdio.h>
#include<stdint.h>
#include<string.h>
#include<stdlib.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <sys/socket.h>



#pragma pack(push, 1)

typedef struct  tagASHEAD {
	uint16_t  HeadID;
	char      Magic;
	uint8_t   Version;
	uint32_t  BodyLen;
	uint32_t  VisitorIP;
	uint16_t  Method;
	uint16_t  PrivLen;
	char      Priv[0];
} ASHEAD, *LPASHEAD;

#pragma pack(pop)

void SendMsg(char * buf, int len) {

        int ret = 0;
        struct sockaddr_in addr;    
        int sock = socket(AF_INET, SOCK_STREAM, 0); 

        bzero(&addr, sizeof(addr));
        addr.sin_family = AF_INET;
        addr.sin_port = htons(8188);
        addr.sin_addr.s_addr = inet_addr("127.0.0.1");

        if(connect(sock, (struct sockaddr *)&addr, sizeof(addr))<0) {
                perror("connect");
                exit(1);
        }   
    

	send(sock, buf, len, 0);

	ret = recv(sock, buf, 1024, 0);
	printf("receive len %d\n", ret);

	printf("Receive %s\n", buf);

}

/*创建授权码*/
#define AU_CREATE_AUCODE_REQ  0x3C01
#define AU_CREATE_AUCODE_RESP 0x3C02

void create_aucode() {
	
	char buf[512];
	char * p;
	int s;
	ASHEAD head = {
		HeadID:     0X0ABA,
		Magic:      '$',
		Version:    0x30,
		VisitorIP:  0,
		Method:     AU_CREATE_AUCODE_REQ,
		PrivLen:    0,
		};
	

	FILE * fp = fopen("./JsonStr.txt", "r+");
	if(fp == NULL) {
		printf("open file failed\n");
		return;
	} 
	
	while(1) {
		p = fgets(buf, sizeof(buf), fp);
		if ( p == NULL ) {
			break;
		}
		p = strtok(buf, "=");
		if( strcmp(p, "CreateAuCode") == 0) {
			p = strtok(NULL, "=");
			break;
		}
	}
	
	printf("I get %s\n", p);
	printf("len is %d\n", sizeof(head));

	char buf2[1024];
	memcpy(buf2, &head, sizeof(head));
	memcpy(buf2+sizeof(head), buf, strlen(buf));

	SendMsg(buf2, strlen(buf)+sizeof(head));
}


/*获取Token */
#define AU_REQUEST_TOKEN_REQ  0x3C03
#define AU_REQUEST_TOKEN_RESP 0x3c04

/*访问AsToken*/
#define AU_ACCESS_TOKEN_REQ   0x3C05
#define AU_ACCESS_TOKEN_RESP  0x3C06

/*刷新AsToken*/
#define AU_REFRESH_TOKEN_REQ  0x3C07
#define AU_REFRESH_TOKEN_RESP 0x3C08

/*检查AsToken*/
#define AU_INSPECT_TOKEN_REQ  0x3C09
#define AU_INSPECT_TOKEN_RESP 0x3C0A

/*注销Token*/
#define AU_DESTROY_TOKEN_REQ   0x3C0B
#define AU_DESTROY_TOKEN_RESP  0x3C0C

/*申请Session*/
#define AU_REQUEST_SESSION_REQ  0x3C0D
#define AU_REQUEST_SESSION_RESP 0x3C0E


int main(int argc, char * argv[]) {

	if (argc < 2) {
		printf("Uage: %s C|G|V|R|I|D|S", argv[0]);
		printf("C: Create Aucode\n");
		printf("G: Get Token\n");
		printf("V: Visit Token\n");
		printf("R: refresh token\n");
		printf("I: Inspect token\n");
		printf("D: Destroy token\n");
		printf("S: Session token\n");
		return;
	}


	switch ( argv[1][0] ) {
		case  'C':
			create_aucode();
			break;
		case  'G':
			printf("G: Get Token\n");
			break;
		case  'V':
			printf("V: Visit Token\n");
			break;
		case  'R':
			printf("R: refresh token\n");
			break;
		case  'I':
			printf("I: Inspect token\n");
			break;
		case  'D':
			printf("D: Destroy token\n");
			break;
		case  'S':
			printf("S: Session token\n");
			break;
		default:
			break;
	}

	return 0;
}
