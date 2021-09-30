from collections import Counter
class Solution:
    def frequencySort(self, s):
        """
        :type s: str
        :rtype: str
        """
        n = len(s)
        sc = dict(Counter(s))
        ll = ['' for i in range(0, n+1)]
        isvisited = set()
        ret = ""
        for key,value in sc.items():
                ll[value] += key * value
        for i in range(n, -1, -1):
            if ll[i] != '':
                ret += ll[i]
        return ret;

def main():
    print("python main function")
    s=Solution()
    s.frequencySort("fdsadfa") 


if __name__ == "__main__":
    main()