//+------------------------------------------------------------------+
//|                                                     JsonPost.mq4 |
//|                        Copyright 2018, MetaQuotes Software Corp. |
//|                                             https://www.mql5.com |
//+------------------------------------------------------------------+
#property copyright "Copyright 2018, MetaQuotes Software Corp."
#property link      "https://www.mql5.com"
#property version   "1.00"
#property strict
#include<JAson.mqh>
void OnStart()
  {
      string hisOrders,resp;
      int histotal = OrdersHistoryTotal();
      for(int i =histotal-1;i>=0;i--)
      {
         string order;
         if(OrderSelect(i,SELECT_BY_POS,MODE_HISTORY))
         {
            if(TimeCurrent() -OrderCloseTime() > 60*60*12) break; 
            string ticket = IntegerToString(OrderTicket());
            string time = TimeToStr(OrderCloseTime()+60*60*6);
            string type = IntegerToString(OrderType());
            double lots = DoubleToStr(OrderLots());
            string symbol = OrderSymbol();
            string price = DoubleToStr(OrderOpenPrice(),MarketInfo(Symbol(),MODE_DIGITS));
            string sl = DoubleToStr(OrderStopLoss(),MarketInfo(Symbol(),MODE_DIGITS));
            string tp = DoubleToStr(OrderTakeProfit(),MarketInfo(Symbol(),MODE_DIGITS));
            string profit = DoubleToStr(OrderProfit()+OrderCommission()+OrderSwap(),2);
            string magic = IntegerToString(OrderMagicNumber());
            order=StringConcatenate(ticket,",",time,",",type,",",lots,",",symbol,",",price,",",sl,",",tp,",",profit,",",magic); 
         }
         StringAdd(hisOrders,order);
         StringAdd(hisOrders,";");

      }
      
      
      resp = StringConcatenate("10000","@","XAUUSD,1,2,0.01","@","usdjpy:11,2,22","@","20000","@","30000","@","50022817");
      
   string cookie=NULL,headers; 
   char post[],result[]; 
   int res; 
//--- to enable access to the server, you should add URL "https://www.google.com/finance" 
//--- in the list of allowed URLs (Main Menu->Tools->Options, tab "Expert Advisors"): 
   StringToCharArray(resp,post);
   string google_url="http://local.com/monit"; 
//--- Reset the last error code 
   ResetLastError(); 
//--- Loading a html page from Google Finance 
   int timeout=5000; //--- Timeout below 1000 (1 sec.) is not enough for slow Internet connection 
   res=WebRequest("GET",google_url,cookie,NULL,timeout,post,0,result,headers); 
//--- Checking errors 
   if(res==-1) 
     { 
      Print("Error in WebRequest. Error code  =",GetLastError()); 
      //--- Perhaps the URL is not listed, display a message about the necessity to add the address 
      MessageBox("Add the address '"+google_url+"' in the list of allowed URLs on tab 'Expert Advisors'","Error",MB_ICONINFORMATION); 
     } 
   else 
     { 
      //--- Load successfully 
      Print(CharArrayToString(result,0,-1,CP_ACP)); 
     } 
  }
//+------------------------------------------------------------------+
