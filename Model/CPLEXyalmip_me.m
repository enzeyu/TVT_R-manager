%% yalmip
% 目标函数 min x1 + 2 x2 + 3 x3 + x4
% 约束条件
% - x1 + x2 + x3 + 10 x4 <= 20
% x1 - 3 x2 + x3 <= 30
% x2 - 3.5 x4 = 0
% 上下陿
% 0 <= x1 <= 40
% 0 <= x2
% 0 <= x3
% 2 <= x4 <= 3
% 整数变量
% x4

clear
clc
tic 

% 设置变量
a = sdpvar(1,2); % a代表单位代价，由服务提供者提供
Num = 2; % 车辆数目
beta = 2; % 浮动奖励公式的浮动参敿
rMax = 100; % 奖励最大值
alpha = 10; % 固定奖励的参数
theta = normrnd(100,20,1,Num); % 初始信誉值

% a的约束和车辆初始信誉值有兿

%% 设约板
c=[a(1)*beta*log(rMax/(a(1)*beta))-theta(1)<= 0;
    -beta*log(rMax/(a(1)*beta))<= 0;
    a(2)*beta*log(rMax/(a(2)*beta))-theta(2) <= 0;
    -beta*log(rMax/(a(2)*beta)) <= 0;
    theta(1)/10>=a(1)>=theta(1)/40;
    theta(2)/10>=a(2)>=theta(2)/40];
objective = a(1)*beta*log(rMax/(a(1)*beta)) - 2*alpha*log2(3) - 2*rMax + a(1)*beta + a(2)*beta*log(rMax/(a(2)*beta)) + a(2)*beta;
%% 设求解器
ops=sdpsettings('verbose', 2,'showprogress',2, 'debug', 1);
sol=optimize(c, objective, ops);
saveampl(c,objective,'mymodel');
%% 分析错误标志
if sol.problem == 0
    disp('success!');
else
    disp('error');
    yalmiperror(sol.problem)
    disp(sol.problem);
end

% 输出的是车辆信誉值
disp(theta(1));
disp(theta(2));
% 输出的是定价a
disp(double(a(1)));
disp(double(a(2)));
% 输出的是ae，即代价
disp(double(a(1))*beta*log(rMax/(double(a(1))*beta)))
disp(double(a(2))*beta*log(rMax/(double(a(2))*beta)))
% 输出的是效用
disp(alpha*log2(3)+rMax*(1-exp(-beta*log(rMax/(double(a(1))*beta))/beta)))
disp(alpha*log2(3)+rMax*(1-exp(-beta*log(rMax/(double(a(2))*beta))/beta)))


    
