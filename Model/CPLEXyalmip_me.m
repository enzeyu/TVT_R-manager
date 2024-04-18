%% yalmip
% Ŀ�꺯�� min x1 + 2 x2 + 3 x3 + x4
% Լ������
% - x1 + x2 + x3 + 10 x4 <= 20
% x1 - 3 x2 + x3 <= 30
% x2 - 3.5 x4 = 0
% �����
% 0 <= x1 <= 40
% 0 <= x2
% 0 <= x3
% 2 <= x4 <= 3
% ��������
% x4

clear
clc
tic 

% ���ñ���
a = sdpvar(1,2); % a����λ���ۣ��ɷ����ṩ���ṩ
Num = 2; % ������Ŀ
beta = 2; % ����������ʽ�ĸ����Δ�
rMax = 100; % �������ֵ
alpha = 10; % �̶������Ĳ���
theta = normrnd(100,20,1,Num); % ��ʼ����ֵ

% a��Լ���ͳ�����ʼ����ֵ�Ѓ�

%% ��Լ��
c=[a(1)*beta*log(rMax/(a(1)*beta))-theta(1)<= 0;
    -beta*log(rMax/(a(1)*beta))<= 0;
    a(2)*beta*log(rMax/(a(2)*beta))-theta(2) <= 0;
    -beta*log(rMax/(a(2)*beta)) <= 0;
    theta(1)/10>=a(1)>=theta(1)/40;
    theta(2)/10>=a(2)>=theta(2)/40];
objective = a(1)*beta*log(rMax/(a(1)*beta)) - 2*alpha*log2(3) - 2*rMax + a(1)*beta + a(2)*beta*log(rMax/(a(2)*beta)) + a(2)*beta;
%% �������
ops=sdpsettings('verbose', 2,'showprogress',2, 'debug', 1);
sol=optimize(c, objective, ops);
saveampl(c,objective,'mymodel');
%% ���������־
if sol.problem == 0
    disp('success!');
else
    disp('error');
    yalmiperror(sol.problem)
    disp(sol.problem);
end

% ������ǳ�������ֵ
disp(theta(1));
disp(theta(2));
% ������Ƕ���a
disp(double(a(1)));
disp(double(a(2)));
% �������ae��������
disp(double(a(1))*beta*log(rMax/(double(a(1))*beta)))
disp(double(a(2))*beta*log(rMax/(double(a(2))*beta)))
% �������Ч��
disp(alpha*log2(3)+rMax*(1-exp(-beta*log(rMax/(double(a(1))*beta))/beta)))
disp(alpha*log2(3)+rMax*(1-exp(-beta*log(rMax/(double(a(2))*beta))/beta)))


    
