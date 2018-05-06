using System.Collections.Generic;
using System.Threading.Tasks;

namespace KpiLex.BusinessLogic.Services.Abstract
{
    public interface IService<TEntity, in TIdentity>
    {
        Task<TEntity> AddAsync(TEntity entity);
        Task UpdateAsync(TEntity entity);
        Task RemoveAsync(TIdentity identity);
        Task<IList<TEntity>> GetAll();
        Task<TEntity> GetById(TIdentity id);
    }
}