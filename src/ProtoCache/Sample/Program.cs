using Microsoft.Extensions.Caching.Distributed;
using ProtoCache;

namespace Sample;

sealed class Program
{
    public static async Task Main()
    {
        var cancellationToken = CancellationToken.None;
        var address = "http://localhost:4000";
        IDistributedCache cache = new ProtoDistributedCache(address);

        var key = "key1";
        var textValue = "sample text";

        var value = System.Text.Encoding.UTF8.GetBytes(textValue);
        await cache.SetAsync(key, value, cancellationToken);

        var result = await cache.GetAsync(key, cancellationToken);
        var textResult = System.Text.Encoding.UTF8.GetString(result);
        if (textResult != textValue)
        {
            throw new Exception("Not equal");
        }
    }
}