Num = 150; % 车辆总数目
Malicious = 75; % 车辆里恶意车辆数目
Prices = sdpvar(1,Num); % 单位代价，对应a
beta = 2; % 浮动奖励公式的浮动参数
rMax = 100; % 奖励最大值
alpha = 10; % 固定奖励的参数
theta = normrnd(400,50,1,Num); % 初始信誉值
theta_up = 10;  %定价的上界
theta_down = 20;    % 定价的下界
Tasknumber = 50; %  任务数目
taxTerm = 2; % 税收周期
allCosts = []; % 记录所有车辆的所有代价
allRewards = [];    % 记录所有车辆的所有奖励
ours_reputationRecords = theta; % 最后的信誉值输出，初始化为theta

for i = 1 : 1: Tasknumber
    [prices,costs,rewards,tax] = oneTask(Prices,Num,beta,rMax,alpha,theta,theta_up,theta_down);
    for j = 1: 1: Num
        % 恶意上报更新信誉值
        if 1 <= j && j <= Malicious 
            theta(j) = theta(j) - costs(j);
 
        % 正常上报更新信誉值
        else
           theta(j) = theta(j) - costs(j) + rewards(j);  % 正常上报信誉值更新

        end

    end
    % 进行周期性税收
    thetaSum = sum(theta);
    if mod(i,taxTerm)==0
    		for j = 1: 1: Num
    				if 1 <= j && j <= Malicious
    						theta(j) = theta(j) - (1/Malicious)*0.5*(-tax);
    				else
    						theta(j) = theta(j) - (theta(j)/thetaSum)*0.25*(-tax);
    				end
    		    if theta(j) < 0
        	    theta(j) = eps;
        		end
    		end
    end
    % 记录信誉值
    %disp(CMRELM_theta);
    ours_reputationRecords = [ours_reputationRecords;theta];
    allCosts = [allCosts,costs];
    allRewards = [allRewards,rewards];
end
