import { HttpError } from '@kubernetes/client-node';
import type { NextApiRequest, NextApiResponse } from 'next';
import { CRDMeta, GetCRD, K8sApi } from '../../../../services/backend/kubernetes';
import { BadRequestResp, InternalErrorResp, JsonResp, UnprocessableResp } from '../../response';

export default async function handler(req: NextApiRequest, resp: NextApiResponse) {
  if (req.method !== 'POST') {
    return BadRequestResp(resp);
  }

  const { kubeconfig } = req.body;
  // console.log(req.body);
  if (kubeconfig === '') {
    return UnprocessableResp('kubeconfig or user', resp);
  }

  const kc = K8sApi(kubeconfig);

  // get user account payment amount

  const user = kc.getCurrentUser();
  if (user === null) {
    return BadRequestResp(resp);
  }

  const account_meta: CRDMeta = {
    group: 'account.sealos.io',
    version: 'v1',
    namespace: 'sealos-system',
    plural: 'accounts'
  };

  type accountStatus = {
    amount: number;
  };

  type amountResp = {
    status: number;
    amount: number;
  };

  try {
    const accountDesc = await GetCRD(kc, account_meta, user.name);
    if (accountDesc !== null && accountDesc.body !== null && accountDesc.body.status !== null) {
      const accountStatus = accountDesc.body.status as accountStatus;
      return JsonResp(
        {
          status: 200,
          amount: accountStatus.amount
        } as amountResp,
        resp
      );
    }
  } catch (err) {
    console.log(err);

    if (err instanceof HttpError) {
      return InternalErrorResp(err.body.message, resp);
    }
  }

  return InternalErrorResp('get amount failed', resp);
}
