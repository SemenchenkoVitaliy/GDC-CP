using System.Threading.Tasks;

namespace KpiLex.Api.Mappers.Abstract
{
    public interface IApiMapper<TIn, TOut>
    {
        TOut Map(TIn model);
    }
}