using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Order.Domain.SeedWork;

namespace Order.Domain.AggregateModels.OrderModels
{
    public class Address:ValueObject
    {
        public string City { get; set; }
        public string Country { get; set; }     
        /// <summary>
        /// ValueObject olarak karşılaşırılmak üzere kullanılır.
        /// </summary>
        /// <returns></returns>
        protected override IEnumerable<object> GetEqualityComponents()
        {
            /*
             * Yield return ifadesi ile iterator’a çağrı yapılan foreach döngüsüne bir eleman
             * döndürülürken yield break ifadesi ile de artık bulunan iterator içerisindeki iterasyonun sona
             * erdiği bilgisi iterator’ı çağıran foreach döngüsüne iletilmekte.
             */
            yield return City;
            yield return Country;
        }
    }
}
