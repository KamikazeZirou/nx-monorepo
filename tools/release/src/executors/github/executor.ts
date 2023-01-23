import { logger } from '@nrwl/devkit';
import { lastValueFrom, of } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { exec } from '../common/exec';
import type { GithubExecutorSchema } from './schema';

export default async function runExecutor({
  tag,
  notes,
  target,
  draft,
  title,
}: GithubExecutorSchema) {
  const createRelease$ = exec('gh', [
    'release',
    'edit',
    tag,
    ...(notes ? ['--notes', notes] : []),
    ...(target ? ['--target', target] : []),
    ...(draft ? ['--draft'] : []),
    ...(title ? ['--title', title] : []),
  ]).pipe(
    map(() => ({ success: true })),
    catchError((response) => {
      logger.info(response);

      return exec('gh', [
        'release',
        'create',
        tag,
        ...(notes ? ['--notes', notes] : []),
        ...(target ? ['--target', target] : []),
        ...(draft ? ['--draft'] : []),
        ...(title ? ['--title', title] : []),
      ]).pipe(
        map(() => ({ success: true })),
        catchError((response) => {
          logger.error(response);
          return of({ success: false });
        })
      );
    })
  );

  return lastValueFrom(createRelease$);
}
