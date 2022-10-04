using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using MediatR;

namespace Order.Domain.Events
{
    public class OrderStartedDomainEvent:INotification
    {
        //order başladığını haber veren bir eventdir
        //mediatr package indirildi
        //eventler application taraınfa handle edilmeli
        //burada sadece order bilgisi yeterli olmaz start eden kullanıcının da bilgileri gerekir

        public string UserName { get; set; }
        public AggregateModels.OrderModels.Order Order { get; set; }

        public OrderStartedDomainEvent(string userName, AggregateModels.OrderModels.Order order)
        {
            UserName = userName;
            Order = order;
        }
    }
}
