using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using MediatR;
using Order.Application.Repository;
using Order.Domain.AggregateModels.BuyerModels;
using Order.Domain.Events;

namespace Order.Application.DomainEventHandlers
{
    public class OrderStartedDomainHandler:INotificationHandler<OrderStartedDomainEvent>
    {
        //meidatr aracılığı ile OrderStartedDomainEvent dan gelenleri handle et diyoruz

        private readonly IBuyerRepository buyerRepository;
        public OrderStartedDomainHandler(IBuyerRepository _buyerRepository)
        {
            this.buyerRepository = _buyerRepository;
        }

        public Task Handle(OrderStartedDomainEvent notification, CancellationToken cancellationToken)
        {
            //order started olduğunda ne olmalı?
            if (notification.Order.UserName == "")
            {
                var buyer = new Buyer(notification.Order.UserName);
                //buyerRepository.Add(buyer)
            }

            return Task.CompletedTask;
        }
    }
}
