% 输入�?系列的参数，获得这个车辆对应的信�?
% a的约束和车辆初始信誉值有�?
% a代表单位代价，Num 车辆数目�?
% beta浮动奖励公式的浮动参数，rMax是奖励最大�?�，alpha是固定奖励的参数，theta是初始信誉�?�，theta_up,theta_down分别代表定价的上下界

function [prices,costs,rewards,tax] = oneTask(Prices,Num,beta,rMax,alpha,theta,theta_up,theta_down)
%% 设约�?
constraints = [];
objective = 0;
for i = 1 : 1: Num
    constraints = constraints+ (Prices(i)*beta*log(rMax/(Prices(i)*beta))-theta(i)<= 0);
    constraints = constraints+ (-beta*log(rMax/(Prices(i)*beta))<= 0);
    constraints = constraints+ (theta(i)/theta_up>=Prices(i)>=theta(i)/theta_down)
    objective = objective+ (Prices(i)*beta*log(rMax/(Prices(i)*beta)) - 2*alpha*log2(3) - 2*rMax + Prices(i)*beta)
end

disp(constraints);
disp(objective);
disp(Prices);

%% 设求解器
ops=sdpsettings('verbose',1);
sol=optimize(constraints, objective, ops);
saveampl(constraints,objective,'mymodel');
%% 分析错误标志
if sol.problem == 0
    disp('success!');
else
    %disp('error');
    yalmiperror(sol.problem);
end

costs = [];
rewards = [];
tax = 0;
prices = Prices;
for i = 1 : 1: Num
    currentCost = double(Prices(i))*beta*log(rMax/(double(Prices(i))*beta));
    costs = [costs,currentCost];
    currentReward = alpha*log2(3)+rMax*(1-exp(-beta*log(rMax/(double(Prices(i))*beta))/beta));
    rewards = [rewards,currentReward];
    tax = tax-currentReward+currentCost;
%     disp("定价�?");
%     disp(double(Prices(i)));
    disp("������?");
    disp(currentCost);
    disp("������?");
    disp(currentReward);
end
    disp(tax);
end
