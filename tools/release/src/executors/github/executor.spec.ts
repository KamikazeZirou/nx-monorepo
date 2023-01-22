import { GithubExecutorSchema } from './schema';
import executor from './executor';

const options: GithubExecutorSchema = {
  tag: 'v1.0.0',
};

describe('Github Executor', () => {
  it('can run', async () => {
    const output = await executor(options);
    expect(output.success).toBe(true);
  });
});
