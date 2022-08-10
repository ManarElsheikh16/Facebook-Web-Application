create database Facebook;

use Facebook;
##################################################################################################################
create table user(
user_ID BIGINT not null primary key,
firstName_user varchar(255) not null,
middleName_user varchar(255) not null,
lastName_user varchar(255) not null,
username varchar(255) not null,
mobile_user varchar(255) not null,
email_user varchar(255) not null,
password_user varchar(255) not null,
registeredAt_user varchar(255) not null,
lastLogin_user varchar(255) not null
);

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(1,"mohamed","soliman","abo haseba","Mohamed Soliman","01028574231","MohamedSoliman@gmail.com","MS#011",'2018-8-18','2020-8-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(2,"nourhan","abd elhafez","mahmoud","Nour Elsheikh","01028574232","NourElsheikh@gmail.com","NE#012",'2017-8-18','2021-8-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(3,"shahd","abd elhafez","mahmoud","Shahd Elsheikh","01028574233","ShahdElsheikh@gmail.com","SE#013",'2016-8-18','2022-8-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(4,"fatma","abd elhafez","mahmoud","Fatma Elsheikh","01028574234","FatmaElsheikh@gmail.com","FE#014",'2015-8-18','2021-10-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(5,"ahmed","mohamed","ali","Ahmed Ali","01028574235","AhmedAli@gmail.com","AA#015",'2014-8-18','2020-4-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(6,"samy","ali","ahmed","Samy Ahmed","01028574236","SamyAhmed@gmail.com","SA#016",'2013-8-18','2021-5-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(7,"shady","abd elmo3m","3aziz","Shady 3aziz","01028574237","Shady3aziz@gmail.com","SA#017",'2012-8-18','2022-6-20');

insert into
user(user_ID,firstName_user,middleName_user,lastName_user,username,mobile_user,email_user,password_user,registeredAt_user,lastLogin_user)
values(8,"eslam","abd elrahman","yousef","eslam yousef","01028574237","EslamYousef@gmail.com","EY#018",'2011-8-18','2021-7-20');

##################################################################################################################
create table user_message(
user_message_ID BIGINT not null primary key,
message_body TEXT not null,
createdAt_userMessage varchar(255) not null,
updatedAt_userMessage varchar(255) not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID)
);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(1,"hi",'2022-8-20','2022-8-21',8);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(2,"how are you",'2022-7-20','2022-7-21',7);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(3,"you are beatiful today",'2022-6-20','2022-6-21',6);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(4,"can you tell me please why dr ahmed need me?",'2022-5-20','2022-5-21',5);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(5,"why dad eat this egg?!",'2022-4-20','2022-4-21',4);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(6,"hope you are well today",'2022-3-20','2022-3-21',3);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(7,"what about going to GYM tomorrow?",'2022-2-20','2022-2-21',2);

insert into
user_message(user_message_ID,message_body,createdAt_userMessage,updatedAt_userMessage,user_ID)
values(8,"taste my food when you back from school",'2022-1-20','2022-1-21',1);

##################################################################################################################
create table user_follower(
user_follower_ID BIGINT not null primary key,
type_userFollower varchar(255) not null,
createdAt_userFollower varchar(255) not null,
updatedAt_userFollower varchar(255) not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID)
);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(1,"like",'2022-1-20','2022-2-20',7);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(2,"dislike",'2022-1-10','2022-5-23',6);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(3,"follow",'2020-2-20','2021-9-20',3);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(4,"like",'2020-2-20','2021-9-20',4);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(5,"dislike",'2019-6-20','2020-2-9',1);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(6,"follow",'2010-3-20','2022-3-20',8);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(7,"follow",'2021-2-10','2022-5-20',5);

insert into
user_follower(user_follower_ID,type_userFollower,createdAt_userFollower,updatedAt_userFollower,user_ID)
values(8,"like",'2018-5-23','2019-6-20',2);

##################################################################################################################
create table user_friend(
user_friend_ID BIGINT not null primary key,
status_userFriend varchar(255) not null,
type_userFriend varchar(255) not null,
notes TEXT not null,
createdAt_userFriend varchar(255) not null,
updatedAt_userFriend varchar(255) not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID)
);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(1,"New","School","No Comment",'2022-1-20','2022-1-22',8);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(2,"New","School","No Comment",'2022-3-10','2022-3-12',7);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(3,"Active","College","No Comment",'2022-2-2','2022-2-8',6);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(4,"Rejected","College","No Comment",'2022-6-20','2022-6-24',5);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(5,"New","College","No Comment",'2022-5-20','2022-5-25',4);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(6,"Active","School","No Comment",'2022-3-25','2022-3-30',3);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(7,"Rejected","College","No Comment",'2022-7-21','2022-7-28',2);

insert into
user_friend(user_friend_ID,status_userFriend,type_userFriend,notes,createdAt_userFriend,updatedAt_userFriend,user_ID)
values(8,"Rejected","School","No Comment",'2022-4-20','2022-4-24',1);

##################################################################################################################
create table user_post(
user_post_ID BIGINT not null primary key,
post_content TEXT not null,
post_language varchar(255) not null,
createdAt_userPost varchar(255) not null,
updatedAt_userPost varchar(255) not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID)
);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(1,'finally i finish my graduation project','English','2010-2-1','2011-2-1',1);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(2,'special day my best friend mariage','English','2012-8-6','2013-8-6',2);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(3,'special day my birth day ','English','2014-4-7','2015-4-7',3);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(4,'i found new community','English','2016-5-5','2017-5-5',4);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(5,'i love my college FCI ','English','2018-11-10','2019-11-10',5);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(6,'bad day my memories','English','2020-7-9','2021-9-9',6);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(7,'buty is pain ','English','2002-2-17','2003-2-17',7);

insert into
user_post(user_post_ID,post_content,post_language,createdAt_userPost,updatedAt_userPost,user_ID)
values(8,'do not cry everything will path','English','2004-3-20','2005-3-20',8);

##################################################################################################################
create table user_memories(
userMemories_ID BIGINT not null primary key,
userMemories_content TEXT not null,
createdAt_userMemories varchar(255) not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID)
);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(1,"a7la day begaad ",'2022-3-20',4);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(2,"graduation Day",'2021-4-20',5);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(3,"finally I finish my course",'2019-3-10',1);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(4,"today manar make my day",'2021-9-6',2);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(5,"Eid Ad7a mobarrk",'2020-3-5',3);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(6,"so bad day i hope it pass",'2021-4-9',2);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(7,"my friend wedding I'm so excited",'2020-1-7',8);

insert into
user_memories(userMemories_ID,userMemories_content,createdAt_userMemories,user_ID)
values(8,"Look at this beautiful flower",'2017-9-9',7);

##################################################################################################################
create table _group(
group_ID BIGINT not null primary key,
createdBy_group BIGINT not null,
updatedBy_group BIGINT not null,
createdAt_group varchar(255) not null,
updatedAt_group varchar(255) not null,
content_group TEXT not null,
status_group varchar(255) not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID)
);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(1,8,7,'2021-8-15','2021-10-21','clothes','Approved',1);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(2,2,4,'2021-2-3','2021-7-28','Cosmetics_store', 'New' ,2);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(3,1,6,'2019-3-13','2020-5-6','education','Active',3);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(4,3,5,'2013-7-1','2021-5-20','games','Blocked',4);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(5,7,2,'2018-2-22','2022-2-1','Courses','New',5);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(6,6,1,'2014-5-2','2020-7-7','Wedding_planer','Active',6);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(7,4,3,'2019-3-22','2022-8-13','Labtop_Market','New',7);

insert into
_group(group_ID,createdBy_group,updatedBy_group,createdAt_group,updatedAt_group,content_group,status_group,user_ID)
values(8,5,7,'2015-5-2','2021-8-5','Travelling','Active',8);

##################################################################################################################
create table group_follower(
groupFollower_ID BIGINT not null primary key,
groupFollower_type varchar(255) not null,
createdAt_groupFollower varchar(255) not null,
updatedAt_groupFollower varchar(255) not null,
group_ID BIGINT not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID),
FOREIGN KEY (group_ID) REFERENCES _group(group_ID)
);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(1,"like",'2015-5-2','2021-8-5',8,1);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(2,"dislike",'2015-7-2','2021-11-5',7,2);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(3,"follow",'2015-7-13','2021-11-20',6,3);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(4,"like",'2015-2-13','2021-10-20',5,4);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(5,"dislike",'2020-2-3','2021-10-10',4,5);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(6,"follow",'2021-3-2','2022-10-15',3,6);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(7,"like",'2015-10-7','2022-12-19',2,7);

insert into
group_follower(groupFollower_ID,groupFollower_type,createdAt_groupFollower,updatedAt_groupFollower,group_ID,user_ID)
values(8,"dislike",'2014-10-17','2022-12-14',1,8);

##################################################################################################################
create table group_post(
groupPost_ID BIGINT not null primary key,
createdAt_groupPost varchar(255) not null,
updatedAt_groupPost varchar(255) not null,
groupPost_content TEXT(255) not null,
group_ID BIGINT not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID),
FOREIGN KEY (group_ID) REFERENCES _group(group_ID)
);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(1,'2020-3-12','2021-3-4','All colors for this dress are available',1,1);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(2,'2012-3-2','2020-5-8','All products are original',2,2);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(3,'2018-6-22','2020-9-4','Anyone who needs to register for a physics course contact us',3,3);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(4,'2021-9-23','2022-3-4','how much this game?',4,4);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(5,'2017-4-15','2021-7-14','i need games for 3 years childern',5,5);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(6,'2020-6-15','2021-9-14','how much to take IT course',6,6);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(7,'2018-3-1','2019-7-13','how can i make a password?',7,7);

insert into
group_post(groupPost_ID,createdAt_groupPost,updatedAt_groupPost,groupPost_content,group_ID,user_ID)
values(8,'2020-5-15','2022-3-14','any weeding planner is available now',8,8);

##################################################################################################################
create table group_meta(
groupMeta_ID BIGINT not null primary key,
groupMeta_content TEXT not null,
group_ID BIGINT not null,
FOREIGN KEY (group_ID) REFERENCES _group(group_ID)
);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(1,'Sadat_region',1);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(2,'6 October',2);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(3,'Sabet_street',3);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(4,'Fryal',4);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(5,'Nemese',5);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(6,'Tawheed_Street',6);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(7,'Assiut_Station',7);

insert into
group_meta(groupMeta_ID,groupMeta_content,group_ID)
values(8,'sheikh_Zayed',8);

##################################################################################################################
create table group_message(
groupMessage_ID BIGINT not null primary key,
message_body TEXT not null,
createdAt_groupMessage varchar(255) not null,
updatedAt_groupMessage varchar(255) not null,
group_ID BIGINT not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID),
FOREIGN KEY (group_ID) REFERENCES _group(group_ID)
);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(1,"how much this blouse?",'2014-10-17','2022-12-14',8,7);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(2,"what size are available?",'2015-10-17','2021-12-14',7,6);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(3,"what color of the dress?",'2014-9-17','2022-12-11',6,5);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(4,"what length of this pants?",'2010-10-17','2022-7-14',5,4);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(5,"is this shimeeze oversized?",'2014-5-7','2022-8-14',4,3);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(6,"what time the store available?",'2020-10-12','2022-12-18',3,2);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(7,"can i see this dress in reality?",'2006-10-7','2022-12-4',2,1);

insert into
group_message(groupMessage_ID,message_body,createdAt_groupMessage,updatedAt_groupMessage,group_ID,user_ID)
values(8,"can i make an order?",'2014-1-10','2022-5-14',1,8);

##################################################################################################################
create table group_member(
groupMember_ID BIGINT not null primary key,
role_ID smallint not null ,
status_groupMember varchar(255) not null,
notes TEXT not null,
createdAt_groupMember varchar(255) not null,
updatedAt_groupMember varchar(255) not null,
group_ID BIGINT not null,
user_ID BIGINT not null,
FOREIGN KEY (user_ID) REFERENCES user(user_ID),
FOREIGN KEY (group_ID) REFERENCES _group(group_ID)
);


insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(1,1,'new','no comments','2003-3-3','2004-3-3',1,1);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(2,2,'rejected','no comments','2005-7-8','2006-7-8',2,2);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(3,3,'active','no comments','2007-8-8','2008-8-15',3,3);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(4,2,'blocked','no comments','2009-5-11','2010-5-11',4,4);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(5,1,'new','no comments','2011-6-3','2012-6-3',5,5);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(6,3,'active','no comments','2013-6-16','2014-6-16',6,6);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(7,2,'new','no comments','2015-3-18','2016-3-18',7,7);

insert into
group_member(groupMember_ID,role_ID,status_groupMember,notes,createdAt_groupMember,updatedAt_groupMember,group_ID,user_ID)
values(8,1,'blocked','no comments','2017-1-16','2018-1-16',8,8);

##################################################################################################################
select * from user;
drop table user;

select * from _group;
drop table _group;

select * from group_member;

select * from group_post;
drop table group_post;



show databases;

select doc_name from doctor order by salary;

update doctor
set address ='Aswan'
where doc_ID =1 or doc_ID=8;